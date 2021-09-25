package config

import (
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client redis.Client
}

func NewRedisClient(host string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
	return &RedisClient{
		client: *client,
	}
}
