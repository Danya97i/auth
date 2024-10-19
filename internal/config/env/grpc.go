package env

import (
	"errors"
	"net"
	"os"

	"github.com/Danya97i/auth/internal/config"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

type grpcConfig struct {
	host string
	port string
}

// NewGrpcConfig создает новый gRPC конфиг
func NewGrpcConfig() (*grpcConfig, error) {
	host := os.Getenv("GRPC_HOST")
	if host == "" {
		return nil, errors.New("GRPC_HOST is not set")
	}
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		return nil, errors.New("GRPC_PORT is not set")
	}
	return &grpcConfig{host, port}, nil
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
