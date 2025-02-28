package infra

import "context"

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Cache --output=../../../mocks
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any) error
}
