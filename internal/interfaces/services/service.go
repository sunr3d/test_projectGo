package services

import (
	"context"
	"time"
)

type InputLink struct {
	Link      string
	FakeLink  string
	EraseTime time.Time
}

type Service interface {
	Create(ctx context.Context, link InputLink) error
}
