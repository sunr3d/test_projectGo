package kafka_impl

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"link_service/internal/config"
	"link_service/internal/interfaces/infra"
)

var _ infra.Broker = (*Kafka)(nil)

type Kafka struct {
	Writer *kafka.Writer
	Conn   *kafka.Conn
	Logger *zap.Logger
}

func New(log *zap.Logger, cfg config.Kafka) (infra.Broker, error) {
	writer := kafka.Writer{
		Addr:     kafka.TCP(cfg.Addr),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	conn, err := kafka.Dial("tcp", cfg.Addr)
	if err != nil {
		return nil, fmt.Errorf("kafka.Dial: %w", err)
	}

	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return nil, fmt.Errorf("conn.SetDeadline: %w", err)
	}

	if _, err = conn.Brokers(); err != nil {
		return nil, fmt.Errorf("conn.Brokers: %w", err)
	}
	log.Info("Connect to Kafka success")

	return &Kafka{Writer: &writer, Conn: conn, Logger: log}, nil
}

func (k *Kafka) AddMsg(ctx context.Context, key []byte, message []byte) error {
	err := k.Writer.WriteMessages(ctx, kafka.Message{Key: key, Value: message})
	if err != nil {
		return fmt.Errorf("k.Writer.WriteMessages: %w", err)
	}
	k.Logger.Debug("Add message to Kafka Success", zap.String("key", string(key)), zap.String("message", string(message)))
	return nil
}

func (k *Kafka) Close() error {
	if err := k.Writer.Close(); err != nil {
		return fmt.Errorf("k.Writer.Close: %w", err)
	}
	return k.Conn.Close()
}
