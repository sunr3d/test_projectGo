package infra

import (
	"context"
	"time"
)

type InputLink struct {
	Link      string
	FakeLink  string
	EraseTime time.Time
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Database --output=../../../mocks
type Database interface {
	Find(ctx context.Context, fakeLink string) (*string, error)
	Create(ctx context.Context, link InputLink) error
}
