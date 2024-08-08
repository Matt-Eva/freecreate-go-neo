package handlers

import (
	"context"
	"freecreate/internal/err"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetWriting(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client) http.HandlerFunc {
 return func (w http.ResponseWriter, r *http.Request){
	getWriting(w, r, ctx, neo,  mongo)
 }
}

func getWriting(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client){
	urlParams := r.URL.Query()
	creatorIds , ok := urlParams["creatorId"]
	if !ok {
		e := err.New("url params does not include creator id")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}

	writingIds, ok := urlParams["writingId"]
	if !ok {
		e := err.New("url params does not include writing id")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}

	if len(creatorIds) != 1 {
		e := err.New("invalid number of creator id params - must be 1 and only 1")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}
	if len(writingIds) != 1  {
		e := err.New("invalid number of writing id params - must be 1 and only 1")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}

	creatorId := creatorIds[0]
	writingId := writingIds[0]

}

func CreateWriting(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleCreateWriting(w, r, ctx, neo, mongo, store)
	}
}

func handleCreateWriting(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, store *redisstore.RedisStore) {

}

func UpdateWriting(w http.ResponseWriter, r *http.Request) {

}

func DeleteWriting(w http.ResponseWriter, r *http.Request) {

}
