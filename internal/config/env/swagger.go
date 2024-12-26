package env

import (
	"fmt"
	"net"
	"os"

	"github.com/Danya97i/auth/internal/config"
)

type swaggerConfig struct {
	host string
	port string
}

// NewSwaggerConfig creates a new swagger config
func NewSwaggerConfig() (config.SwaggerConfig, error) {
	host := os.Getenv("SWAGGER_HOST")
	if len(host) == 0 {
		return nil, fmt.Errorf("SWAGGER_HOST is not set")
	}

	port := os.Getenv("SWAGGER_PORT")
	if len(port) == 0 {
		return nil, fmt.Errorf("SWAGGER_PORT is not set")
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}

func (c *swaggerConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
