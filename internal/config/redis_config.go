package config

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"os"

	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	fmt.Println("connecting redis")
	address := os.Getenv("REDIS_ADDRESS")
	pwd := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: pwd,
		DB:       0,
	})
	fmt.Println("redis connected")
	return client
}

func InitRedisSessionStore(ctx context.Context, redis *redis.Client)(*redisstore.RedisStore, err.Error){
	fmt.Println("initializing session store")
	store, rErr := redisstore.NewRedisStore(ctx, redis)
	if rErr != nil {
		return store, err.NewFromErr(rErr)
	}
	fmt.Println("session store initialized")
	return store, err.Error{}
}
