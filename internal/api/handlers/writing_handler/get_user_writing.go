package writing_handler

import (
	"context"
	"freecreate/internal/api/middleware"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

func GetUserWriting(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc{
return func (w http.ResponseWriter, r *http.Request){
getUserWriting(w, r, ctx, neo, store)
}
}

func getUserWriting(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore){
	user, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	
}