package cache

import (
	"context"
	"time"

	"github.com/s3rzh/go-grpc-user-service/pkg/cache/redis"
)

type Cache interface {
	Set(context.Context, string, interface{}, time.Duration) (string, error)
	Get(context.Context, string) (string, error)
	Delete(context.Context, string) (string, error)
}

func NewCache(cfg redis.Config) (Cache, error) {
	return redis.NewRedisCache(cfg)
}
