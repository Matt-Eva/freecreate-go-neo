package handlers

import (
	"context"
	"encoding/json"
	"freecreate/internal/api/middleware"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

func GetCreator(w http.ResponseWriter, r *http.Request) {

}

func CreateCreator(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		createCreator(w, r, ctx, neo, store)
	}
}



func createCreator(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore){
	user, uErr := middleware.AuthenticateUser(r, store)
	if uErr.E != nil{
		http.Error(w, uErr.E.Error(), http.StatusUnauthorized)
	}

	body := r.Body
	postedCreator := models.PostedCreator{}
	json.NewDecoder(body).Decode(&postedCreator)

	creator, mErr := models.GenerateCreator(user.Uid, postedCreator)
	if mErr.E != nil {
		http.Error(w, mErr.E.Error(), http.StatusUnprocessableEntity)
	}

	createdCreator, cErr := queries.CreateCreator(ctx, neo, user, creator)
	if cErr.E != nil {
		http.Error(w, cErr.E.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(createdCreator)
}

func UpdateCreator(w http.ResponseWriter, r *http.Request) {

}

func DeleteCreator(w http.ResponseWriter, r *http.Request) {

}
