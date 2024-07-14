package routes

import (
	"context"
	"freecreate/internal/api/handlers"
	"freecreate/internal/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
)

func CreateRoutes(ctx context.Context, mongo string, neo neo4j.DriverWithContext, redis *redis.Client) error {
	router := mux.NewRouter()

	// TEST ENDPOINTS
	// =====================

	router.HandleFunc("/api", middleware.AddDrivers(handlers.TestHandler, neo, mongo, redis, ctx)).Methods("GET")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(handlers.TestCachePostHandler, redis, ctx)).Methods("POST")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(handlers.TestCacheGetHandler, redis, ctx)).Methods("GET")

	// APPLICATION ENDPOINTS
	// =====================

	// SEARCH ROUTES
	// no name, no tags, time frame != mostRecent, query cache (Redis? Mongo?)
	router.HandleFunc("/api/search/cache", middleware.AddRedisDriver(handlers.SearchCacheHandler, redis, ctx)).Methods("GET")

	// timeFrame == mostRecent || name || tags
	router.HandleFunc("/api/search/standard", middleware.AddNeoDriver(handlers.SearchStandardHandler, neo)).Methods("GET")

	// time frame == mostRecent - query neo current year, order by date, not rank
	router.HandleFunc("/api/search/most-recent", middleware.AddNeoDriver(handlers.SearchMostRecentHandler, neo)).Methods("GET")

	// name || tags && timeFrame == allTime - query neo allTime db, order by absolute rank
	router.HandleFunc("/api/search/all-time", middleware.AddNeoDriver(handlers.SearchAllTimeHandler, neo)).Methods("GET")

	router.HandleFunc("/api/default-content", middleware.AddMongoDriver(handlers.DefaultContentHandler, mongo)).Methods("GET")

	// AUTH ROUTES
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/logout", handlers.Logout).Methods("DELETE")

	// USER ROUTES
	router.HandleFunc("/api/user", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user", handlers.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/user", handlers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/user/reading-history", handlers.GetReadingHistory).Methods("GET")
	router.HandleFunc("/api/user/reading-history", handlers.AddToReadingHistory).Methods("POST")
	router.HandleFunc("/api/user/reading-history", handlers.UpdateReadingHistory).Methods("PATCH")
	router.HandleFunc("/api/user/reading-history", handlers.RemoveFromReadingHistory).Methods("DELETE")
	router.HandleFunc("/api/user/likes", handlers.GetLikes).Methods("GET")
	router.HandleFunc("/api/user/likes", handlers.CreateLike).Methods("POST")
	router.HandleFunc("/api/user/likes", handlers.DeleteLike).Methods("DELETE")
	router.HandleFunc("/api/user/library", handlers.GetLibrary).Methods("GET")
	router.HandleFunc("/api/user/library", handlers.AddToLibrary).Methods("POST")
	router.HandleFunc("/api/user/library", handlers.RemoveFromLibrary).Methods("DELETE")
	router.HandleFunc("/api/user/reading-list", handlers.GetReadingList).Methods("GET")
	router.HandleFunc("/api/user/reading-list", handlers.AddToReadingList).Methods("POST")
	router.HandleFunc("/api/user/reading-list", handlers.RemoveFromReadingList).Methods("DELETE")
	router.HandleFunc("/api/user/bookshelf", handlers.GetBookshelf).Methods("GET")
	router.HandleFunc("/api/user/bookshelf/bookshelf", handlers.CreateBookshelf).Methods("POST")
	router.HandleFunc("/api/user/bookshelf/item", handlers.AddToBookshelf).Methods("POST")
	router.HandleFunc("/api/user/bookshelf/bookshelf", handlers.DeleteBookshelf).Methods("DELETE")
	router.HandleFunc("/api/user/bookshelf/item", handlers.RemoveFromBookshelf).Methods("DELETE")
	router.HandleFunc("/api/user/subscriptions", handlers.GetSubscriptions).Methods("GET")
	router.HandleFunc("/api/user/subscriptions", handlers.CreateSubscription).Methods("POST")
	router.HandleFunc("/api/user/subscriptions", handlers.DeleteSubscription).Methods("DELETE")
	router.HandleFunc("/api/user/following", handlers.GetFollowing).Methods("GET")
	router.HandleFunc("/api/user/following", handlers.Follow).Methods("POST")
	router.HandleFunc("/api/user/following", handlers.Unfollow).Methods("DELETE")

	// CREATOR ROUTES
	// router.HandleFunc("/api/creator").Methods("GET")
	// router.HandleFunc("/api/creator").Methods("POST")
	// router.HandleFunc("/api/creator").Methods("PATCH")
	// router.HandleFunc("/api/creator").Methods("DELETE")

	// WRITING ROUTES
	// router.HandleFunc("/api/short-story").Methods("GET")
	// router.HandleFunc("/api/short-story").Methods("POST")
	// router.HandleFunc("/api/short-story").Methods("PATCH")
	// router.HandleFunc("/api/short-story").Methods("DELETE")
	// router.HandleFunc("/api/novellete").Methods("GET")
	// router.HandleFunc("/api/novellete").Methods("POST")
	// router.HandleFunc("/api/novellete").Methods("PATCH")
	// router.HandleFunc("/api/novellete").Methods("DELETE")
	// router.HandleFunc("/api/novella").Methods("GET")
	// router.HandleFunc("/api/novella").Methods("POST")
	// router.HandleFunc("/api/novella").Methods("PATCH")
	// router.HandleFunc("/api/novella").Methods("DELETE")
	// router.HandleFunc("/api/novel").Methods("GET")
	// router.HandleFunc("/api/novel").Methods("POST")
	// router.HandleFunc("/api/novel").Methods("PATCH")
	// router.HandleFunc("/api/novel").Methods("DELETE")
	// router.HandleFunc("/api/chapter").Methods("GET")
	// router.HandleFunc("/api/chapter").Methods("POST")
	// router.HandleFunc("/api/chapter").Methods("PATCH")
	// router.HandleFunc("/api/chapter").Methods("DELETE")
	// router.HandleFunc("/api/collection").Methods("GET")
	// router.HandleFunc("/api/collection/collection").Methods("POST")
	// router.HandleFunc("/api/collection/item").Methods("POST")
	// router.HandleFunc("/api/collection").Methods("PATCH")
	// router.HandleFunc("/api/collection/collection").Methods("DELETE")
	// router.HandleFunc("/api/collection/item").Methods("DELETE")
	// router.HandleFunc("/api/universe").Methods("GET")
	// router.HandleFunc("/api/universe").Methods("POST")
	// router.HandleFunc("/api/universe").Methods("PATCH")
	// router.HandleFunc("/api/universe").Methods("DELETE")

	// DONATION ROUTES
	// router.HandleFunc("/api/donation/given").Methods("GET")
	// router.HandleFunc("/api/donation/received").Methods("GET")
	// router.HandleFunc("/api/donation").Methods("POST")
	// router.HandleFunc("/api/donation").Methods("PATCH")
	// router.HandleFunc("/api/donation").Methods("DELETE")

	// timeFrame == previous year - query neo specific year, order by rank && rel_rank - DEPRECATED
	// router.HandleFunc("/api/search/year", middleware.AddNeoDriver(handlers.SearchYearHandler, neo)).Methods("GET")

	return http.ListenAndServe(":8080", router)
}
