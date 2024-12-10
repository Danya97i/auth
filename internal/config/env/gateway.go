package env

import (
	"fmt"
	"net"
	"os"

	"github.com/Danya97i/auth/internal/config"
)

type gatewayConfig struct {
	host string
	port string
}

func NewGatewayConfig() (config.GatewayConfig, error) {
	host := os.Getenv("GATEWAY_HOST")
	if len(host) == 0 {
		return nil, fmt.Errorf("GATEWAY_HOST is not set")
	}

	port := os.Getenv("GATEWAY_PORT")
	if len(port) == 0 {
		return nil, fmt.Errorf("GATEWAY_PORT is not set")
	}

	return &gatewayConfig{
		host: host,
		port: port,
	}, nil
}

func (c *gatewayConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
