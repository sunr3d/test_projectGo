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

func (s *service) Create(ctx context.Context, link services.InputLink) error {
	err := s.repo.Find(ctx, link.FakeLink)
	if err != nil {
		return errors.New("ALREADY IN USE")
	}
	err = s.repo.Create(ctx, infra.InputLink(link))
	return err
}
