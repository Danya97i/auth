package env

import (
	"errors"
	"os"

	"github.com/Danya97i/auth/internal/config"
)

var _ config.PGConfig = (*pgConfig)(nil)

type pgConfig struct {
	dsn string
}

// NewPgConfig создает новый pgConfig
func NewPgConfig() (*pgConfig, error) {
	dsn := os.Getenv("PG_DSN")
	if dsn == "" {
		return nil, errors.New("empty postgres dsn string")
	}
	return &pgConfig{dsn: dsn}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
