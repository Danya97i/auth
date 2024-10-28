package app

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/Danya97i/auth/internal/config"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

// App - собираем все зависимости приложения
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
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
	return a.runGrpcServer()
}

// initDeps - инициализирует зависимости
func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcServer,
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
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)
	pb.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserServer(ctx))
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
