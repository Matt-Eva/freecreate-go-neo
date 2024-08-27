package chapters

import (
	"context"
	"freecreate/internal/middleware"
	"net/http"

	"github.com/rbcervilla/redisstore/v9"
	"go.mongodb.org/mongo-driver/mongo"
)


func UpdateChapterNumberHandler(ctx context.Context, mongo *mongo.Client, store *redisstore.RedisStore)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		handleUpdateChapterNumber(w, r, ctx, mongo, store)
	}
}

func handleUpdateChapterNumber(w http.ResponseWriter, r *http.Request, ctx context.Context, mongo *mongo.Client, store *redisstore.RedisStore){
	_, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
	}
}