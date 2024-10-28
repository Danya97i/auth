package pg

import (
	"context"

	"github.com/Danya97i/auth/internal/client/db"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type pgClient struct {
	masterDBC db.DB
}

// NewPGClient создает новый pg клиент
func NewPGClient(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}
	return &pgClient{
		masterDBC: &pg{dbc},
	}, nil
}

// DB возвращает master dbc
func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

// Close закрывает соединение с master dbc
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}
	return nil
}
