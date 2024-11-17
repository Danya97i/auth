package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Handler – функция для выполнения в транзакции
type Handler func(ctx context.Context) error

// TxManager – менеджер транзакций
type TxManager interface {
	ReadCommited(ctx context.Context, f Handler) error
}

// Transactor – инициатор транзакции
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Query – запрос
type Query struct {
	Name     string
	RawQuery string
}

// Client – клиент к базе данных
type Client interface {
	DB() DB
	Close() error
}

// SQLExecer – исполнитель запросов к базе данных
type SQLExecer interface {
	QueryExecer
	NamedExecer
}

// QueryExecer – исполнитель запросов к базе данных
type QueryExecer interface {
	ExecContext(ctx context.Context, query Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryRowContext(ctx context.Context, Query Query, args ...interface{}) pgx.Row
}

// NamedExecer – запросы со сканированием результата
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// Pinger – пингер
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB – база данных
type DB interface {
	SQLExecer
	NamedExecer
	Transactor
	Pinger
	Close()
}
