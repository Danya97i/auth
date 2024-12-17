package app

import (
	"context"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/Danya97i/platform_common/pkg/closer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/Danya97i/auth/internal/config"
	"github.com/Danya97i/auth/internal/interceptor"
	pb "github.com/Danya97i/auth/pkg/user_v1"

	// register statik
	_ "github.com/Danya97i/auth/statik"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

// App - собираем все зависимости приложения
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	gatewayServer   *http.Server
	swaggerServer   *http.Server
}

// NewApp -создает экземпляр приложения
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Run - запускает приложение
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()

		err := a.runGrpcServer()
		if err != nil {
			log.Fatalf("failed to run grpc server: %s", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runGatewayServer()
		if err != nil {
			log.Fatalf("failed to run gateway server: %s", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runSwaggerServer()
		if err != nil {
			log.Fatalf("failed to run swagger server: %s", err)
		}
	}()

	wg.Wait()

	return nil
}

// initDeps - инициализирует зависимости
func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcServer,
		a.initGatewayServer,
		a.initSwaggerServer,
	}
	for _, f := range initFuncs {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

// initConfig - инициализирует конфиг
func (a *App) initConfig(_ context.Context) error {
	if err := config.Load(configPath); err != nil {
		return err
	}
	return nil
}

// initServiceProvider - инициализирует сервис-провайдер
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// initGrpcServer - инициализирует gRPC сервер
func (a *App) initGrpcServer(ctx context.Context) error {
	creds, err := credentials.NewServerTLSFromFile("service.pem", "service.key")
	if err != nil {
		return err
	}

	a.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)
	reflection.Register(a.grpcServer)
	pb.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserServer(ctx))
	return nil
}

func (a *App) initGatewayServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterUserV1HandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.gatewayServer = &http.Server{
		Addr:              a.serviceProvider.GatewayConfig().Address(),
		Handler:           corsMiddleware.Handler(mux),
		ReadHeaderTimeout: time.Second * 3,
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	mux.HandleFunc("/api.swagger.json", serveSwaggerFile("/api.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:              a.serviceProvider.SwaggerConfig().Address(),
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 3,
	}

	return nil
}

// runGrpcServer - запускает gRPC сервер
func (a *App) runGrpcServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())
	lis, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	return a.grpcServer.Serve(lis)
}

func (a *App) runGatewayServer() error {
	log.Printf("gateway server is running on %s", a.serviceProvider.GatewayConfig().Address())

	err := a.gatewayServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	log.Printf("swagger server is running on %s", a.serviceProvider.SwaggerConfig().Address())

	err := a.swaggerServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		statikFS, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("open swagger file %s", path)

		file, err := statikFS.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer func() { _ = file.Close() }()

		log.Printf("read swagger file %s", path)

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("write swagger file: %s", path)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("served swagger file %s", path)
	}
}
