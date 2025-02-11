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

type Database interface {
	Find(ctx context.Context, fakeLink string) (string, error)
	Create(ctx context.Context, link InputLink) (int, error)
}
