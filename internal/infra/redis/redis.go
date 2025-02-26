package redis_impl

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"link_service/internal/config"
	"link_service/internal/interfaces/infra"
	"time"
)

var _ infra.Cache = (*RedisDB)(nil)

type RedisDB struct {
	Logger *zap.Logger
	Client *redis.Client
}

func New(lg *zap.Logger, cfg config.Redis) (*RedisDB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	lg.Info("Connect to Redis database success")

	return &RedisDB{Logger: lg, Client: client}, nil
}

func (r *RedisDB) Close() error {
	return r.Client.Close()
}

func (r *RedisDB) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisDB) Set(ctx context.Context, key string, value any) error {
	ttl := 24 * time.Hour
	return r.Client.Set(ctx, key, value, ttl).Err()
}
