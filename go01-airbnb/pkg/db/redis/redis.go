package redis

import (
	"context"
	"fmt"
	"go01-airbnb/config"

	"github.com/go-redis/redis/v8"
)

var (
	defaultRedisMaxActive = 0 // 0 unlimitted max active connection
	defaultRedisMaxIdle   = 10
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password, // no password set
		DB:           cfg.Redis.DB,       // use default DB
		PoolSize:     defaultRedisMaxActive,
		MinIdleConns: defaultRedisMaxIdle,
	})

	// Ping to test connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
