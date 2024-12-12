package producer

import (
	"context"

	"github.com/IBM/sarama"
)

type producer struct {
	syncP sarama.SyncProducer
	topic string
}

func NewProducer(syncP sarama.SyncProducer, topic string) *producer {
	return &producer{
		syncP: syncP,
		topic: topic,
	}
}

func (p *producer) SendMessage(ctx context.Context, message []byte) error {
	msg := sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := p.syncP.SendMessage(&msg)
	return err
}

func (p *producer) Close() error {
	return p.syncP.Close()
}
