package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	RedisClient *redis.Client
}

func NewRedis(redisClient *redis.Client) RedisRepository {
	return &Redis{RedisClient: redisClient}
}

type RedisRepository interface {
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

func (r Redis) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	status := r.RedisClient.Set(ctx, key, value, ttl)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func (r Redis) Get(ctx context.Context, key string) (string, error) {
	value := r.RedisClient.Get(ctx, key)
	if value.Err() != nil {
		return "", value.Err()
	}
	return value.Val(), nil
}
