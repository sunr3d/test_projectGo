package link_service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"link_service/internal/interfaces/infra"
	"link_service/internal/interfaces/services"
)

var _ services.Service = (*service)(nil)

type service struct {
	logger zap.Logger
	repo   infra.Database
}

func NewService(logger zap.Logger, repo infra.Database) services.Service {
	return &service{logger: logger, repo: repo}
}

func (s *service) Create(ctx context.Context, link services.InputLink) (int, error) {
	var id int
	linkFound, err := s.repo.Find(ctx, link.FakeLink)
	if err != nil {
		return id, err
	}
	if linkFound != "" {
		return id, errors.New("link already exists")
	}

	id, err = s.repo.Create(ctx, infra.InputLink(link))

	return id, nil
}

func (s *service) Find(ctx context.Context, fakeLink string) (string, error) {
	link, err := s.repo.Find(ctx, fakeLink)
	if err != nil {
		return "", err
	}
	return link, nil
}
