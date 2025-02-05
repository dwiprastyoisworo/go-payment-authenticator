package config

import "github.com/redis/go-redis/v9"

func (config *AppConfig) RedisInit() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	return rdb
}
