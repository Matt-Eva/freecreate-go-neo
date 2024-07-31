package routes

import (
	"context"
	"freecreate/internal/api/handlers"
	"freecreate/internal/api/middleware"
	"freecreate/internal/api/test_handlers"
	"freecreate/internal/err"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRoutes(ctx context.Context,  neo neo4j.DriverWithContext, mongo *mongo.Client, redis *redis.Client, sessionStore *redisstore.RedisStore) err.Error {
	router := mux.NewRouter()

	// TEST ENDPOINTS
	// =====================

	// router.HandleFunc("/api", middleware.AddDrivers(test_handlers.TestHandler, neo, mongo, redis, ctx)).Methods("GET")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(test_handlers.TestCachePostHandler, redis, ctx)).Methods("POST")
	router.HandleFunc("/api/test-cache", middleware.AddRedisDriver(test_handlers.TestCacheGetHandler, redis, ctx)).Methods("GET")
	router.HandleFunc("/api/master-user", test_handlers.HandleMasterUser(ctx, neo, sessionStore)).Methods("GET")

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

	// timeFrame == previous year - query neo specific year, order by rank && rel_rank
	// router.HandleFunc("/api/search/year", middleware.AddNeoDriver(handlers.SearchYearHandler, neo)).Methods("GET")

	// router.HandleFunc("/api/default-content", middleware.AddMongoDriver(handlers.DefaultContentHandler, mongo)).Methods("GET")

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
	router.HandleFunc("/api/creator", handlers.GetCreator).Methods("GET")
	router.HandleFunc("/api/creator", handlers.CreateCreator).Methods("POST")
	router.HandleFunc("/api/creator", handlers.UpdateCreator).Methods("PATCH")
	router.HandleFunc("/api/creator", handlers.DeleteCreator).Methods("DELETE")

	// WRITING ROUTES
	router.HandleFunc("/api/writing", handlers.GetWriting).Methods("GET")
	router.HandleFunc("/api/writing", handlers.CreateWriting(ctx, neo, mongo, sessionStore)).Methods("POST")
	router.HandleFunc("/api/writing", handlers.UpdateWriting).Methods("PATCH")
	router.HandleFunc("/api/writing", handlers.DeleteWriting).Methods("DELETE")
	router.HandleFunc("/api/draft", handlers.GetDraft).Methods("GET")
	router.HandleFunc("/api/draft", handlers.CreateDraft).Methods("POST")
	router.HandleFunc("/api/draft", handlers.UpdateDraft).Methods("PATCH")
	router.HandleFunc("/api/draft", handlers.DeleteDraft).Methods("DELETE")
	router.HandleFunc("/api/short-story", handlers.GetShortStory).Methods("GET")
	router.HandleFunc("/api/short-story", handlers.CreateShortStory).Methods("POST")
	router.HandleFunc("/api/short-story", handlers.UpdateShortStory).Methods("PATCH")
	router.HandleFunc("/api/short-story", handlers.DeleteShortStory).Methods("DELETE")
	router.HandleFunc("/api/novellete", handlers.GetNovelette).Methods("GET")
	router.HandleFunc("/api/novellete", handlers.CreateNovelette).Methods("POST")
	router.HandleFunc("/api/novellete", handlers.UpdateNovelette).Methods("PATCH")
	router.HandleFunc("/api/novellete", handlers.DeleteNovelette).Methods("DELETE")
	router.HandleFunc("/api/novella", handlers.GetNovella).Methods("GET")
	router.HandleFunc("/api/novella", handlers.CreateNovella).Methods("POST")
	router.HandleFunc("/api/novella", handlers.UpdateNovella).Methods("PATCH")
	router.HandleFunc("/api/novella", handlers.DeleteNovella).Methods("DELETE")
	router.HandleFunc("/api/novel", handlers.GetNovel).Methods("GET")
	router.HandleFunc("/api/novel", handlers.CreateNovel).Methods("POST")
	router.HandleFunc("/api/novel", handlers.UpdateNovel).Methods("PATCH")
	router.HandleFunc("/api/novel", handlers.DeleteNovel).Methods("DELETE")
	router.HandleFunc("/api/chapter", handlers.GetChapter).Methods("GET")
	router.HandleFunc("/api/chapter", handlers.CreateChapter).Methods("POST")
	router.HandleFunc("/api/chapter", handlers.UpdateChapter).Methods("PATCH")
	router.HandleFunc("/api/chapter", handlers.DeleteChapter).Methods("DELETE")
	router.HandleFunc("/api/collection", handlers.GetCollection).Methods("GET")
	router.HandleFunc("/api/collection/collection", handlers.CreateCollection).Methods("POST")
	router.HandleFunc("/api/collection/item", handlers.AddToCollection).Methods("POST")
	router.HandleFunc("/api/collection", handlers.UpdateCollection).Methods("PATCH")
	router.HandleFunc("/api/collection/collection", handlers.DeleteCollection).Methods("DELETE")
	router.HandleFunc("/api/collection/item", handlers.RemoveFromCollection).Methods("DELETE")
	router.HandleFunc("/api/universe", handlers.GetUniverse).Methods("GET")
	router.HandleFunc("/api/universe/universe", handlers.CreateUniverse).Methods("POST")
	router.HandleFunc("/api/universe/item", handlers.AddToUniverse).Methods("POST")
	router.HandleFunc("/api/universe", handlers.UpdateUniverse).Methods("PATCH")
	router.HandleFunc("/api/universe/universe", handlers.DeleteUniverse).Methods("DELETE")
	router.HandleFunc("/api/universe/item", handlers.RemoveFromUniverse).Methods("DELETE")

	// DONATION ROUTES
	router.HandleFunc("/api/donation/given", handlers.GetGivenDonations).Methods("GET")
	router.HandleFunc("/api/donation/received", handlers.GetReceivedDonations).Methods("GET")
	router.HandleFunc("/api/donation", handlers.CreateDonation).Methods("POST")

	hErr := http.ListenAndServe(":8080", router)
	if hErr != nil {
		e := err.NewFromErr(hErr)
		return e
	}

	return err.Error{}
}
