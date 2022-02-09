package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(cfg Config) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisCache{client: client}, nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, exp time.Duration) (string, error) {
	_, err := r.client.Set(ctx, key, value, exp).Result()
	if err != nil {
		return "", err
	}
	return key, nil
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil

}

func (r *RedisCache) Delete(ctx context.Context, key string) (string, error) {
	_, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return key, nil
}
