package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Danya97i/auth/internal/config"
	"github.com/Danya97i/auth/internal/config/env"
	pb "github.com/Danya97i/auth/pkg/user_v1"
	"github.com/jackc/pgx/v4/pgxpool"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	// Загружаем конфиг
	if err := config.Load(configPath); err != nil {
		log.Fatal(err)
	}

	// инициализируем grpc конфиг
	grpcConig, err := env.NewGrpcConfig()
	if err != nil {
		log.Fatal(err)
	}

	// инициализируем postgres конфиг
	pgConfig, err := env.NewPgConfig()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", grpcConig.Address())
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterUserV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
