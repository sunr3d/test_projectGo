package link_service_impl

import (
	"go.uber.org/zap"
	"link_service/internal/interfaces/infra"
)

type service struct {
	logger *zap.Logger
	repo   infra.Database
	cache  infra.Cache
}
