package handlers

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"freecreate/internal/utils"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

type ReturnUser struct {
	Uid string `json:"uid"`
	DisplayName          string `json:"displayName"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	BirthDay             int    `json:"birthday"`
	BirthYear            int    `json:"birthYear"`
	BirthMonth           int    `json:"birthMonth"`
	ProfilePic           string `json:"profilePic"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
return func (w http.ResponseWriter, r *http.Request){
	createUser(w, r, ctx, neo, store)
}
}

type PostedUser struct {
	DisplayName          string `json:"displayName"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	BirthDay             int    `json:"birthday"`
	BirthYear            int    `json:"birthYear"`
	BirthMonth           int    `json:"birthMonth"`
	ProfilePic           string `json:"profilePic"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func createUser(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	var postedUser PostedUser
	if e := json.NewDecoder(r.Body).Decode(&postedUser); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	var postedUserModel models.PostedUser
	if e := utils.StructToStruct(postedUser, postedUserModel); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	userModel, mErr := models.GenerateUser(postedUserModel)
	if mErr.E != nil {
		mErr.Log()
		http.Error(w, mErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	createdUser, cErr := queries.CreateUser(ctx, neo, userModel)
	if cErr.E != nil {
		cErr.Log()
		http.Error(w, cErr.E.Error(), http.StatusInternalServerError)
		return
	}

	var returnUser ReturnUser
	if e := utils.StructToStruct(createdUser, returnUser); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e:= json.NewEncoder(w).Encode(returnUser); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

}



func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
