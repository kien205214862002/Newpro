package cache

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}

type redisCache struct {
	store *cache.Cache
}

func NewRedisCache(redisClient *redis.Client) *redisCache {
	store := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return &redisCache{store}
}

func (r *redisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.store.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   0,
	})
}

func (r *redisCache) Get(ctx context.Context, key string, value interface{}) error {
	return r.store.Get(ctx, key, value)
}

func (r *redisCache) Delete(ctx context.Context, key string) error {
	return r.store.Delete(ctx, key)
}
