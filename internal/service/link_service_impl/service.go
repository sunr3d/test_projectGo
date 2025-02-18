package link_service_impl

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"link_service/internal/interfaces/infra"
	"link_service/internal/interfaces/services"
)

var _ services.Service = (*service)(nil)

type service struct {
	logger *zap.Logger
	repo   infra.Database
}

func New(logger *zap.Logger, repo infra.Database) services.Service {
	return &service{logger: logger, repo: repo}
}

func (s *service) Create(ctx context.Context, link services.InputLink) error {
	linkFound, err := s.repo.Find(ctx, link.FakeLink)
	if err != nil {
		return err
	}
	if linkFound != nil {
		return LinkAlreadyExists
	}
	err = s.repo.Create(ctx, infra.InputLink(link))

	return nil
}

func (s *service) Find(ctx context.Context, fakeLink string) (string, error) {
	link, err := s.repo.Find(ctx, fakeLink)
	if err != nil {
		return "", fmt.Errorf("failed to find link: %w", err)
	}
	if link == nil {
		return "", LinkNotFound
	}
	return *link, nil
}
