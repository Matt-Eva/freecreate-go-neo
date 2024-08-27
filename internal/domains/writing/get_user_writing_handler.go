package writing

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"freecreate/internal/middleware"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

func GetUserWritingHandler(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getUserWriting(w, r, ctx, neo, store)
	}
}

func getUserWriting(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	user, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	writing, status, qErr := GetUserWriting(ctx, neo, user.Uid)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
		return
	}

	returnedWriting := make([]ReturnedWriting, 0)

	for _, writ := range writing {
		retWrit := &ReturnedWriting{}

		retWrit.Uid = writ.Uid
		retWrit.Title = writ.Title
		retWrit.Description = writ.Description
		retWrit.Author = writ.Author
		retWrit.Font = writ.Font
		retWrit.UniqueAuthorName = writ.UniqueAuthorName
		retWrit.CreatorId = writ.CreatorId
		retWrit.Published = writ.Published
		retWrit.Genres = writ.Genres
		retWrit.Tags = writ.Tags

		returnedWriting = append(returnedWriting, (*retWrit))
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(returnedWriting); e != nil {
		newFromE := err.NewFromErr(e)
		newFromE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
