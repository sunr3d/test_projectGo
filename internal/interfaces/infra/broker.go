package infra

import "context"

type Broker interface {
	Add(ctx context.Context, topic string, message []byte) error
}
