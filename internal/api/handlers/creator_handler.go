package handlers

import (
	"context"
	"encoding/json"
	"freecreate/internal/api/middleware"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"freecreate/internal/utils"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

func GetCreator(w http.ResponseWriter, r *http.Request) {

}

func CreateCreator(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createCreator(w, r, ctx, neo, store)
	}
}

type PostedCreator struct {
	CreatorName string `json:"creatorName"`
	CreatorId   string `json:"creatorId"`
	About       string `json:"about"`
}

type ResponseCreator struct {
	CreatorName string `json:"creatorName"`
	CreatorId   string `json:"creatorId"`
	About       string `json:"about"`
	Uid			string `json:"uid"`
}

func createCreator(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	user, uErr := middleware.AuthenticateUser(r, store)
	if uErr.E != nil {
		http.Error(w, uErr.E.Error(), http.StatusUnauthorized)
		return
	}

	body := r.Body
	var postedCreator PostedCreator
	if e := json.NewDecoder(body).Decode(&postedCreator); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, newE.E.Error(), http.StatusInternalServerError)
		return
	}

	var newCreator models.NewCreator
	if e := utils.StructToStruct(postedCreator, newCreator); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	creatorModel, mErr := models.GenerateCreator(user.Uid, newCreator)
	if mErr.E != nil {
		mErr.Log()
		http.Error(w, mErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	var queryCreator queries.NewCreator
	if e := utils.StructToStruct(creatorModel, queryCreator); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	createdCreator, cErr := queries.CreateCreator(ctx, neo, user, creatorModel)
	if cErr.E != nil {
		cErr.Log()
		http.Error(w, cErr.E.Error(), http.StatusInternalServerError)
		return
	}

	responseCreator := ResponseCreator{}
	if e := utils.StructToStruct(createdCreator, responseCreator); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(responseCreator); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, newE.E.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateCreator(w http.ResponseWriter, r *http.Request) {

}

func DeleteCreator(w http.ResponseWriter, r *http.Request) {

}
