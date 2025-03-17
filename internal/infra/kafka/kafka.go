package kafka_impl

import (
	"context"
	"link_service/internal/interfaces/infra"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var _ infra.Broker = (*Kafka)(nil)

type Kafka struct {
	Writer *kafka.Writer
	Logger *zap.Logger
}

func New(log *zap.Logger, port string) infra.Broker {
	writer := kafka.Writer{
		Addr:     kafka.TCP("localhost:" + port),
		Topic:    "link_service",
		Balancer: &kafka.LeastBytes{},
	}

	return &Kafka{Writer: &writer, Logger: log}
}

func (k *Kafka) Add(ctx context.Context, topic string, message []byte) error {
	return k.Writer.WriteMessages(ctx, kafka.Message{Topic: topic, Value: message})
}

func (k *Kafka) Close() error {
	return k.Writer.Close()
}
