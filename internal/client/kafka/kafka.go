package kafka

import "context"

// Producer is a kafka producer
type Producer interface {
	SendMessage(ctx context.Context, message []byte) error
	Close() error
}
