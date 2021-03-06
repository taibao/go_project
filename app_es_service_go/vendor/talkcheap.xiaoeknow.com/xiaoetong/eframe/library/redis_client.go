package library

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(ctx context.Context, conf *RedisConfig) (*RedisClient, error) {
	addr := fmt.Sprintf("%s:%d", conf.Addr, conf.Port)
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	if _, err := cli.Ping(ctx).Result(); err != nil {
		err = fmt.Errorf("redis client:[%s] error:%w", conf.ConnectionName, err)
		return nil, err
	}
	return &RedisClient{cli}, nil
}
