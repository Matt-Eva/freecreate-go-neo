package middleware

import (
	"context"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddDrivers(handler func(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext, mongo *mongo.Client, redis *redis.Client, ctx context.Context), neo neo4j.DriverWithContext, mongo *mongo.Client, redis *redis.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, mongo, redis, ctx)
	}
}

func AddMongoDriver(handler func(w http.ResponseWriter, r *http.Request, mongo *mongo.Client), mongo *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, mongo)
	}
}

func AddNeoDriver(handler func(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext), neo neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo)
	}
}

func AddRedisDriver(handler func(w http.ResponseWriter, r *http.Request, redis *redis.Client, ctx context.Context), redis *redis.Client, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, redis, ctx)
	}
}

func AddNeoAndMongo(handler func(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext, mongo *mongo.Client), neo neo4j.DriverWithContext, mongo *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, mongo)
	}
}

func AddNeoAndRedis(handler func(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext, redis *redis.Client), neo neo4j.DriverWithContext, redis *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, redis)
	}
}

func AddMongoAndRedis(handler func(w http.ResponseWriter, r *http.Request, mongo *mongo.Client, redis *redis.Client), mongo *mongo.Client, redis *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, mongo, redis)
	}
}
