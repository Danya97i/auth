package kafka

import "context"

type Producer interface {
	SendMessage(ctx context.Context, message []byte) error
	Close() error
}
