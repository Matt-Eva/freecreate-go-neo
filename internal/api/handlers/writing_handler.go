package handlers

import (
	"context"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetWriting(w http.ResponseWriter, r *http.Request) {

}

func CreateWriting(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, store *redisstore.RedisStore) http.HandlerFunc {
return func (w http.ResponseWriter, r *http.Request) {
	handleCreateWriting(w, r, ctx, neo, mongo, store)
}
}

func handleCreateWriting(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, store *redisstore.RedisStore){

}

func UpdateWriting(w http.ResponseWriter, r *http.Request) {

}

func DeleteWriting(w http.ResponseWriter, r *http.Request) {

}
