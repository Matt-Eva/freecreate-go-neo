package users

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"freecreate/internal/middleware"
	"freecreate/internal/utils"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)


func GetUser(store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getUser(w, r, store)
	}
}

func getUser(w http.ResponseWriter, r *http.Request, store *redisstore.RedisStore) {
	authenticatedUser, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	var returnUser ReturnUser
	if e := utils.StructToStruct(authenticatedUser, &returnUser); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(returnUser); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusUnauthorized)
	}
}

func CreateUserHandler(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createUserHandler(w, r, ctx, neo, store)
	}
}

type PostedUser struct {
	UniqueName           string `json:"uniqueName"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	BirthDay             int    `json:"birthDay"`
	BirthYear            int    `json:"birthYear"`
	BirthMonth           int    `json:"birthMonth"`
	ProfilePic           string `json:"profilePic"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	var postedUser PostedUser
	if e := json.NewDecoder(r.Body).Decode(&postedUser); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	userModel, mErr := GenerateUser(postedUser)
	if mErr.E != nil {
		mErr.Log()
		http.Error(w, mErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	createdUser, cErr := CreateUserQuery(ctx, neo, userModel)
	if cErr.E != nil {
		cErr.Log()
		http.Error(w, cErr.E.Error(), http.StatusInternalServerError)
		return
	}

	var authenticatedUser middleware.AuthenticatedUser
	if e := utils.StructToStruct(createdUser, &authenticatedUser); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	sErr := middleware.CreateUserSession(w, r, store, authenticatedUser)
	if sErr.E != nil {
		sErr.Log()
		http.Error(w, sErr.E.Error(), http.StatusInternalServerError)
		return
	}

	var returnUser ReturnUser
	if e := utils.StructToStruct(createdUser, &returnUser); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(returnUser); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUser(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateUser(w, r, ctx, neo, store)
	}
}

type PatchedUser struct {
	UniqueName string `json:"uniqueName"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	BirthDay   int    `json:"birthDay"`
	BirthYear  int    `json:"birthYear"`
	BirthMonth int    `json:"birthMonth"`
}

func updateUser(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	authenticatedUser, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	var patchedUser PatchedUser
	if e := json.NewDecoder(r.Body).Decode(&patchedUser); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	updatedUserModel, vErr := GenerateUpdatedUserInfo(patchedUser)
	if vErr.E != nil {
		vErr.Log()
		http.Error(w, vErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	updatedUser, uErr := UpdateUserInfo(ctx, neo, authenticatedUser.Uid, updatedUserModel)
	if uErr.E != nil {
		uErr.Log()
		http.Error(w, uErr.E.Error(), http.StatusInternalServerError)
		return
	}

	var newAuthUser middleware.AuthenticatedUser
	if e := utils.StructToStruct(updatedUser, &newAuthUser); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	sErr := middleware.CreateUserSession(w, r, store, newAuthUser)
	if sErr.E != nil {
		sErr.Log()
		http.Error(w, sErr.E.Error(), http.StatusInternalServerError)
		return
	}

	var returnUser ReturnUser
	if e := utils.StructToStruct(updatedUser, &returnUser); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(returnUser); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteUser(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteUser(w, r, ctx, neo, store)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	authenticatedUser, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	dErr := DeleteUserQuery(ctx, neo, authenticatedUser.Uid)
	if dErr.E != nil {
		dErr.Log()
		http.Error(w, dErr.E.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdatePassword() {

}
