package config

import (
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/library"
)

var DefaultRedisConfig *library.RedisConfig
var ConfhubRedisConfig *library.RedisConfig

func InitRedisConfig(env *library.Env) {
	DefaultRedisConfig = &library.RedisConfig{
		ConnectionName: "redis-default",
		Addr:     env.GetString("REDIS_DEFAULT_RW_HOST"),
		Port:     env.GetInt("REDIS_DEFAULT_RW_PORT"),
		Password: env.GetString("REDIS_DEFAULT_RW_PASSWORD"),
		DB:       env.GetInt("REDIS_DEFAULT_DATABASE"),
		PoolSize: env.GetInt("REDIS_DEFAULT_POOL_SIZE"),
	}

	ConfhubRedisConfig = &library.RedisConfig{
		ConnectionName: "redis-confhub",
		Addr:     env.GetString("REDIS_CONFHUB_RW_HOST"),
		Port:     env.GetInt("REDIS_CONFHUB_RW_PORT"),
		Password: env.GetString("REDIS_CONFHUB_RW_PASSWORD"),
		DB:       env.GetInt("REDIS_CONFHUB_DATABASE"),
		PoolSize: env.GetInt("REDIS_CONFHUB_POOL_SIZE"),
	}
}
