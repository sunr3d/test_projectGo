package services

import (
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Service --output=../../../mocks
type Service interface {
	Create(ctx context.Context, link InputLink) error
	Find(ctx context.Context, fakeLink string) (string, error)
}
