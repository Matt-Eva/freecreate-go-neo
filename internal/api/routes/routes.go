package routes

import (
	"freecreate/internal/api/handlers"
	"freecreate/internal/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoutes(neo, mongo, redis string) error {
	router := mux.NewRouter()

	router.HandleFunc("/api", middleware.AddDrivers(handlers.TestHandler, neo, mongo, redis)).Methods("GET")
	router.HandleFunc("/api/search/most-recent", middleware.AddNeoDriver(handlers.SearchMostRecentHandler, neo)).Methods("GET")
	router.HandleFunc("/api/search/cache", middleware.AddRedisDriver(handlers.SearchCacheHandler, redis)).Methods("GET")

	return http.ListenAndServe(":8080", router)
}
