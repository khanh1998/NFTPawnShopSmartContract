package config

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/model"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Panic(err)
	}
	log.Println("establish a connection to redis server successfully")
	return &RedisClient{
		client: *client,
	}
}

func (r *RedisClient) Put(key string, value interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jsonStr, err := json.Marshal(value)
	if err != nil {
		log.Panic(err)
	}
	err = r.client.Set(ctx, key, jsonStr, 10*time.Minute).Err()
	if err != nil {
		return false
	}
	return true
}

func (r *RedisClient) Get(key string, res *model.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jsonStr, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	err = json.Unmarshal([]byte(jsonStr), res)
	if err != nil {
		return false, err
	}
	return true, nil
}
