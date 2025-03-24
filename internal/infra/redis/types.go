package redis_impl

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type RedisDB struct {
	Logger *zap.Logger
	Client *redis.Client
}
