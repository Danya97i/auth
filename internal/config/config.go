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

// GatewayConfig - конфигурация для Gateway
type GatewayConfig interface {
	Address() string
}

// SwaggerConfig - конфигурация для Swagger
type SwaggerConfig interface {
	Address() string
}

// KafkaConfig - конфигурация для Kafka
type KafkaConfig interface {
	Hosts() string
	UserTopic() string
	MaxRetryCount() int
}

// Load - загружает конфигурацию из файла .env
func Load(path string) error {
	return godotenv.Load(path)
}
