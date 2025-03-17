package services

import (
	"context"
	"github.com/segmentio/kafka-go"
)

//go:generate go run github.com/vektra/mockery/v2@v2.53.2 --name=Service --output=../../../mocks
type Service interface {
	Create(ctx context.Context, link InputLink) error
	Find(ctx context.Context, fakeLink string) (string, error)
	AddMessage(ctx context.Context, msg kafka.Message) error
}
