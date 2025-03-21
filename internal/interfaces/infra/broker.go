package infra

import "context"

//go:generate go run github.com/vektra/mockery/v2@v2.53.2 --name=Broker --output=../../../mocks
type Broker interface {
	Add(ctx context.Context, topic string, key []byte, message []byte) error
}
