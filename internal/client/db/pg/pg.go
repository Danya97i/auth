package pg

import (
	"context"

	"github.com/Danya97i/auth/internal/client/db"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type key string

// TxKey is the context key for the transaction
const TxKey key = "tx"

type pg struct {
	dbc *pgxpool.Pool
}

// NewDB создает новое подключение к базе данных
func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{
		dbc: dbc,
	}
}

// ScanOneContext выполняет запрос и возвращает результат
func (p *pg) ScanOneContext(ctx context.Context, dest any, query db.Query, args ...any) error {
	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, rows)
}

// ScanOneContext выполняет запрос и возвращает результат
func (p *pg) ScanAllContext(ctx context.Context, dest any, query db.Query, args ...any) error {
	rows, err := p.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(dest, rows)
}

// QueryContext выполняет запрос и возвращает результат
func (p *pg) QueryContext(ctx context.Context, query db.Query, args ...any) (pgx.Rows, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, query.RawQuery, args...)
	}
	return p.dbc.Query(ctx, query.RawQuery, args...)
}

// QueryRowContext выполняет запрос и возвращает результат
func (p *pg) QueryRowContext(ctx context.Context, query db.Query, args ...any) pgx.Row {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, query.RawQuery, args...)
	}
	return p.dbc.QueryRow(ctx, query.RawQuery, args...)
}

// ExecContext выполняет запрос
func (p *pg) ExecContext(ctx context.Context, query db.Query, args ...any) (pgconn.CommandTag, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, query.RawQuery, args...)
	}
	return p.dbc.Exec(ctx, query.RawQuery, args...)
}

// BeginTx создает транзакцию
func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

// Ping проверяет соединение
func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

// Close закрывает подключение
func (p *pg) Close() {
	p.dbc.Close()
}

// MakeContextTx создает кладет транзакцию в контекст
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}
