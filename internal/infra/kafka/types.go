package kafka_impl

import (
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Kafka struct {
	Writer *kafka.Writer
	Conn   *kafka.Conn
	Logger *zap.Logger
}
