package middleware

import (
	"net/http"
)

func AddDrivers(handler func(w http.ResponseWriter, r *http.Request, neo, mongo, redis string), neo, mongo, redis string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, mongo, redis)
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

func AddRedisDriver(handler func(w http.ResponseWriter, r *http.Request, redis string), redis string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, redis)
	}
}

func AddNeoAndMongo(handler func(w http.ResponseWriter, r *http.Request, neo, mongo string), neo, mongo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, mongo)
	}
}

func AddNeoAndRedis(handler func(w http.ResponseWriter, r *http.Request, neo, redis string), neo, redis string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, redis)
	}
}

func AddMongoAndRedis(handler func(w http.ResponseWriter, r *http.Request, mongo, redis string), mongo, redis string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, mongo, redis)
	}
}
