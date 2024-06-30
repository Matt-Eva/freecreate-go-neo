package routes

import (
	"freecreate/internal/api/handlers"
	"freecreate/internal/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func CreateRoutes(neo, mongo string, redis *redis.Client) error {
	router := mux.NewRouter()

	// TEST ENDPOINTS
	// =====================

	router.HandleFunc("/api", middleware.AddDrivers(handlers.TestHandler, neo, mongo, redis)).Methods("GET")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(handlers.TestCacheHandler, redis)).Methods("POST")

	// APPLICATION ENDPOINTS
	// =====================

	// no name, no tags, time frame != mostRecent, query Redis cache
	router.HandleFunc("/api/search/cache", middleware.AddRedisDriver(handlers.SearchCacheHandler, redis)).Methods("GET")

	// time frame == mostRecent - query neo current year, order by date, not rank
	router.HandleFunc("/api/search/most-recent", middleware.AddNeoDriver(handlers.SearchMostRecentHandler, neo)).Methods("GET")

	// name || tags && timeFrame == allTime - query neo allTime db, order by absolute rank
	router.HandleFunc("/api/search/all-time", middleware.AddNeoDriver(handlers.SearchAllTimeHandler, neo)).Methods("GET")

	// timeFrame == previous year - query neo specific year, order by rank && rel_rank
	router.HandleFunc("/api/search/year", middleware.AddNeoDriver(handlers.SearchYearHandler, neo)).Methods("GET")

	// timeFrame !== mostRecent || allTime, name || tags - query neo current year, order by rank && rel_rank
	router.HandleFunc("/api/search/standard", middleware.AddNeoDriver(handlers.SearchStandardHandler, neo)).Methods("GET")

	return http.ListenAndServe(":8080", router)
}
