package redis_impl_test

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"link_service/internal/infra/redis"
)

func TestRedisDB_Get(t *testing.T) {
	logger := zap.NewNop()
	client, mock := redismock.NewClientMock()
	redisDB := &redis_impl.RedisDB{Logger: logger, Client: client}

	key := "test_key"
	expectedValue := "test_value"

	mock.ExpectGet(key).SetVal(expectedValue)

	value, err := redisDB.Get(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisDB_Get_Error(t *testing.T) {
	logger := zap.NewNop()
	client, mock := redismock.NewClientMock()
	redisDB := &redis_impl.RedisDB{Logger: logger, Client: client}

	key := "test_key"
	expectedError := redis.Nil

	mock.ExpectGet(key).SetErr(expectedError)

	value, err := redisDB.Get(context.Background(), key)
	assert.Error(t, err)
	assert.Equal(t, "", value)
	assert.Equal(t, expectedError, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisDB_Set(t *testing.T) {
	logger := zap.NewNop()
	client, mock := redismock.NewClientMock()
	redisDB := &redis_impl.RedisDB{Logger: logger, Client: client}

	key := "test_key"
	value := "test_value"
	ttl := 24 * time.Hour

	mock.ExpectSet(key, value, ttl).SetVal("OK")

	err := redisDB.Set(context.Background(), key, value)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRedisDB_Set_Error(t *testing.T) {
	logger := zap.NewNop()
	client, mock := redismock.NewClientMock()
	redisDB := &redis_impl.RedisDB{Logger: logger, Client: client}

	key := "test_key"
	value := "test_value"
	ttl := 24 * time.Hour
	expectedError := errors.New("set error")

	mock.ExpectSet(key, value, ttl).SetErr(expectedError)

	err := redisDB.Set(context.Background(), key, value)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
