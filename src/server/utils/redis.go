package utils

import (
	"context"
	"opengin/server/config"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisHandler struct {
	Client *redis.Client
}

func NewRedisHandler() *RedisHandler {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Settings.Redis.Address,
		Password: config.Settings.Redis.Password,
		DB:       config.Settings.Redis.Db,
		PoolSize: config.Settings.Redis.PoolSize,
	})
	return &RedisHandler{Client: client}
}

func (r *RedisHandler) Get(key string) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	value, err := r.Client.Get(ctx, key).Result()

	if err == nil {
		return value
	}

	return nil
}

func (r *RedisHandler) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	err := r.Client.Set(ctx, key, value, expiration).Err()

	return err
}
