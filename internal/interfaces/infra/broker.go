package infra

import "context"

//go:generate go run github.com/vektra/mockery/v2@v2.53.2 --name=Broker --output=../../../mocks
type Broker interface {
	AddMsg(ctx context.Context, key []byte, message []byte) error
}
