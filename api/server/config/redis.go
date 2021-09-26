package config

import (
	"context"
	"log"
	"time"

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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := client.Ping(ctx).Err(); err != nil {
		log.Panic(err)
	}
	log.Println("establish a connection to redis server successfully")
	return &RedisClient{
		client: *client,
	}
}
