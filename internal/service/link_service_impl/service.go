package link_service_impl

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"link_service/internal/interfaces/infra"
	"link_service/internal/interfaces/services"
)

var _ services.Service = (*service)(nil)

func New(logger *zap.Logger, repo infra.Database, cache infra.Cache) services.Service {
	return &service{logger: logger, repo: repo, cache: cache}
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
	// Ищем в кэше
	cachedLink, err := s.cache.Get(ctx, fakeLink)
	if err == nil && cachedLink != "" {
		//fmt.Println("Link from cache: ", cachedLink) // TODO: DELETE DEBUG LINE
		return cachedLink, nil
	}

	// Если в кэше нет, ищем в БД
	link, err := s.repo.Find(ctx, fakeLink)
	if err != nil {
		return "", fmt.Errorf("failed to find link: %w", err)
	}
	if link == nil {
		return "", LinkNotFound
	}
	//fmt.Println("Link from DB: ", *link) // TODO: DELETE DEBUG LINE

	// Сохраняем в кэш
	if err = s.cache.Set(ctx, fakeLink, *link); err != nil {
		return "", fmt.Errorf("failed to cache link: %w", err)
	}
	//fmt.Println("Saved to cache!", *link) // TODO: DELETE DEBUG LINE

	return *link, nil
}
