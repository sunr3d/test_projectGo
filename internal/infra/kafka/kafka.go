package kafka_impl

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"link_service/internal/interfaces/infra"
)

var _ infra.Broker = (*Kafka)(nil)

func New(log *zap.Logger, port string) (infra.Broker, error) {
	writer := kafka.Writer{
		Addr:     kafka.TCP("localhost:" + port),
		Topic:    "link_service",
		Balancer: &kafka.LeastBytes{},
	}

	conn, err := kafka.Dial("tcp", "localhost:"+port)
	if err != nil {
		return nil, fmt.Errorf("kafka_impl.New: %w", err)
	}

	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return nil, fmt.Errorf("kafka_impl.New: %w", err)
	}

	if _, err = conn.Brokers(); err != nil {
		return nil, fmt.Errorf("kafka_impl.New: %w", err)
	}
	log.Info("Connect to Kafka success")

	return &Kafka{Writer: &writer, Conn: conn, Logger: log}, nil
}

func (k *Kafka) AddMsg(ctx context.Context, key []byte, message []byte) error {
	err := k.Writer.WriteMessages(ctx, kafka.Message{Key: key, Value: message})
	if err != nil {
		return fmt.Errorf("kafka_impl.AddMsg: %w", err)
	}
	k.Logger.Debug("Add message to Kafka Success", zap.String("key", string(key)), zap.String("message", string(message)))
	return nil
}

func (k *Kafka) Close() error {
	if err := k.Writer.Close(); err != nil {
		return fmt.Errorf("kafka_impl.Close: %w", err)
	}
	return k.Conn.Close()
}
