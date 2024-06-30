package middleware

import (
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func AddDrivers(handler func(w http.ResponseWriter, r *http.Request, neo, mongo string, redis *redis.Client, ctx context.Context), neo, mongo string, redis *redis.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, mongo, redis, ctx)
	}
}

func AddMongoDriver(handler func(w http.ResponseWriter, r *http.Request, mongo string), mongo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, mongo)
	}
}

func AddNeoDriver(handler func(w http.ResponseWriter, r *http.Request, neo string), neo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo)
	}
}

func AddRedisDriver(handler func(w http.ResponseWriter, r *http.Request, redis *redis.Client, ctx context.Context), redis *redis.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, redis, ctx)
	}
}

func AddNeoAndMongo(handler func(w http.ResponseWriter, r *http.Request, neo, mongo string), neo, mongo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, mongo)
	}
}

func AddNeoAndRedis(handler func(w http.ResponseWriter, r *http.Request, neo string, redis *redis.Client), neo string, redis *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, redis)
	}
}

func AddMongoAndRedis(handler func(w http.ResponseWriter, r *http.Request, mongo string, redis *redis.Client), mongo string, redis *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, mongo, redis)
	}
}
