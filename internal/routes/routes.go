package routes

import (
	"context"

	"freecreate/internal/domains/auth"
	"freecreate/internal/domains/chapters"
	"freecreate/internal/domains/creators"
	search_handler "freecreate/internal/domains/search"
	"freecreate/internal/domains/users"
	"freecreate/internal/domains/writing"
	"freecreate/internal/err"
	"freecreate/internal/middleware"
	"freecreate/internal/test_handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRoutes(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, redis *redis.Client, store *redisstore.RedisStore) err.Error {
	router := mux.NewRouter()

	// TEST ENDPOINTS
	// =====================

	// router.HandleFunc("/api", middleware.AddDrivers(test_handlers.TestHandler, neo, mongo, redis, ctx)).Methods("GET")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(test_handlers.TestCachePostHandler, redis, ctx)).Methods("POST")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(test_handlers.TestCacheGetHandler, redis, ctx)).Methods("GET")
	router.HandleFunc("/api/master-user", test_handlers.HandleMasterUser(ctx, neo, store)).Methods("GET")

	// APPLICATION ENDPOINTS
	// =====================

	// SEARCH ROUTES
	// no name, no tags, time frame != mostRecent, query cache (Redis? Mongo?)
	router.HandleFunc("/api/search/cache", middleware.AddRedisDriver(search_handler.SearchCacheHandler, redis, ctx)).Methods("GET")

	// timeFrame == mostRecent || name || tags
	router.HandleFunc("/api/search/standard", middleware.AddNeoDriver(search_handler.SearchStandardHandler, neo)).Methods("GET")

	// time frame == mostRecent - query neo current year, order by date, not rank
	router.HandleFunc("/api/search/most-recent", middleware.AddNeoDriver(search_handler.SearchMostRecentHandler, neo)).Methods("GET")

	// name || tags && timeFrame == allTime - query neo allTime db, order by absolute rank
	router.HandleFunc("/api/search/all-time", middleware.AddNeoDriver(search_handler.SearchAllTimeHandler, neo)).Methods("GET")

	// timeFrame == previous year - query neo specific year, order by rank && rel_rank
	// router.HandleFunc("/api/search/year", middleware.AddNeoDriver(handlers.SearchYearHandler, neo)).Methods("GET")

	// router.HandleFunc("/api/default-content", middleware.AddMongoDriver(handlers.DefaultContentHandler, mongo)).Methods("GET")

	// AUTH ROUTES
	router.HandleFunc("/api/login", auth.Login).Methods("POST")
	router.HandleFunc("/api/logout", auth.Logout(store)).Methods("DELETE")

	// USER ROUTES
	router.HandleFunc("/api/user", users.GetUser(store)).Methods("GET")
	router.HandleFunc("/api/user", users.CreateUserHandler(ctx, neo, store)).Methods("POST")
	router.HandleFunc("/api/user", users.UpdateUser(ctx, neo, store)).Methods("PATCH")
	router.HandleFunc("/api/user", users.DeleteUser(ctx, neo, store)).Methods("DELETE")

	// router.HandleFunc("/api/user/likes", like_handler.GetLikes).Methods("GET")
	// router.HandleFunc("/api/user/likes", like_handler.CreateLike).Methods("POST")
	// router.HandleFunc("/api/user/likes", like_handler.DeleteLike).Methods("DELETE")

	// router.HandleFunc("/api/user/library", library_handler.GetLibrary).Methods("GET")
	// router.HandleFunc("/api/user/library", library_handler.AddToLibrary).Methods("POST")
	// router.HandleFunc("/api/user/library", library_handler.RemoveFromLibrary).Methods("DELETE")

	// router.HandleFunc("/api/user/reading-list", reading_list_handler.GetReadingList).Methods("GET")
	// router.HandleFunc("/api/user/reading-list", reading_list_handler.AddToReadingList).Methods("POST")
	// router.HandleFunc("/api/user/reading-list", reading_list_handler.RemoveFromReadingList).Methods("DELETE")

	// router.HandleFunc("/api/user/bookshelf", bookshelf_handler.GetBookshelf).Methods("GET")
	// router.HandleFunc("/api/user/bookshelf/bookshelf", bookshelf_handler.CreateBookshelf).Methods("POST")
	// router.HandleFunc("/api/user/bookshelf/item", bookshelf_handler.AddToBookshelf).Methods("POST")
	// router.HandleFunc("/api/user/bookshelf/bookshelf", bookshelf_handler.DeleteBookshelf).Methods("DELETE")
	// router.HandleFunc("/api/user/bookshelf/item", bookshelf_handler.RemoveFromBookshelf).Methods("DELETE")

	// router.HandleFunc("/api/user/subscriptions", subscription_handler.GetSubscriptions).Methods("GET")
	// router.HandleFunc("/api/user/subscriptions", subscription_handler.CreateSubscription).Methods("POST")
	// router.HandleFunc("/api/user/subscriptions", subscription_handler.DeleteSubscription).Methods("DELETE")

	// router.HandleFunc("/api/user/following", follow_handler.GetFollowing).Methods("GET")
	// router.HandleFunc("/api/user/following", follow_handler.Follow).Methods("POST")
	// router.HandleFunc("/api/user/following", follow_handler.Unfollow).Methods("DELETE")

	// CREATOR ROUTES - creator handler
	router.HandleFunc("/api/creator", creators.GetCreator(ctx, neo)).Methods("GET")
	router.HandleFunc("/api/creator", creators.CreateCreator(ctx, neo, store)).Methods("POST")
	router.HandleFunc("/api/creator", creators.UpdateCreator(ctx, neo, store)).Methods("PATCH")
	router.HandleFunc("/api/creator", creators.DeleteCreator).Methods("DELETE")
	router.HandleFunc("/api/user/creators", creators.GetUserCreators(ctx, neo, store)).Methods("GET")

	// WRITING ROUTES
	router.HandleFunc("/api/writing", writing.GetWritingHandler(ctx, neo)).Methods("GET")
	router.HandleFunc("/api/writing", writing.CreateWritingHandler(ctx, neo, store)).Methods("POST")
	router.HandleFunc("/api/writing", writing.UpdateWritingHandler(ctx, neo, store)).Methods("PATCH")
	router.HandleFunc("/api/writing", writing.DeleteWriting()).Methods("DELETE")
	router.HandleFunc("/api/writing/user", writing.GetUserWritingHandler(ctx, neo, store))

	// CHAPTER ROUTES
	// router.HandleFunc("/api/chapter", chapters.GetChapterHandler(ctx, neo, mongo, store)).Methods("GET")
	router.HandleFunc("/api/chapter", chapters.CreateChapterHandler(ctx, neo, mongo, store)).Methods("POST")
	router.HandleFunc("/api/chapter/number", chapters.UpdateChapterNumberHandler(ctx, mongo, store)).Methods("PATCH")
	router.HandleFunc("/api/chapters", chapters.GetChaptersHandler(ctx, mongo)).Methods("GET")

	// DONATION ROUTES
	// router.HandleFunc("/api/donation/given", donation_handler.GetGivenDonations).Methods("GET")
	// router.HandleFunc("/api/donation/received", donation_handler.GetReceivedDonations).Methods("GET")
	// router.HandleFunc("/api/donation", donation_handler.CreateDonation).Methods("POST")

	hErr := http.ListenAndServe(":8080", router)
	if hErr != nil {
		e := err.NewFromErr(hErr)
		return e
	}

	return err.Error{}
}
