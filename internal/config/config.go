package config

import (
	"time"

	"github.com/joho/godotenv"
)

// GRPCConfig - конфигурация для gRPC
type GRPCConfig interface {
	Address() string
}

// PGConfig - конфигурация для PostgreSQL
type PGConfig interface {
	DSN() string
}

// RedisConfig - конфигурация для Redis
type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

// Load - загружает конфигурацию из файла .env
func Load(path string) error {
	return godotenv.Load(path)
}
