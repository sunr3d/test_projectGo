package link_service_impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
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
		return ErrLinkAlreadyExists
	}
	err = s.repo.Create(ctx, infra.InputLink(link))
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Find(ctx context.Context, fakeLink string) (string, error) {
	// Ищем в кэше
	cachedLink, err := s.cache.Get(ctx, fakeLink)
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			s.logger.Warn("cache.Get err", zap.Error(err))
		}
	}
	if cachedLink != "" {
		s.logger.Debug("cache.Get", zap.String("link", cachedLink))
		return cachedLink, nil
	}

	// Если в кэше нет, ищем в БД
	link, err := s.repo.Find(ctx, fakeLink)
	if err != nil {
		return "", fmt.Errorf("repo.Find err: %w", err)
	}
	if link == nil {
		return "", ErrLinkNotFound
	}
	s.logger.Debug("repo.Find", zap.String("link", *link))

	// Сохраняем в кэш отдельным процессом
	go func() {
		if err = s.cache.Set(context.WithoutCancel(ctx), fakeLink, *link); err != nil {
			s.logger.Error("cache.Set err:", zap.Error(err))
		}
	}()

	return *link, nil
}
