package creator_handler

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

type ResponseCreator struct {
	Name      string `json:"name"`
	UniqueName string `json:"uniqueName"`
	About     string `json:"about"`
	Uid       string `json:"uid"`
}

func (r ResponseCreator) validateResponseCreator() err.Error {
	if r.Name == "" {
		return err.New("response creator name cannot be empty")
	}
	if r.Uid == "" {
		return err.New("response creator Uid cannot be empty")
	}
	if r.UniqueName == "" {
		return err.New("response creator UniqueName cannot be empty")
	}

	return err.Error{}
}

// GET CREATOR
func GetCreator(ctx context.Context, neo neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getCreator(w, r, ctx, neo)
	}
}

func getCreator(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext) {
	urlParams := r.URL.Query()
	creatorIds, ok := urlParams["creatorId"]
	if !ok {
		e := err.New("url params does not include creatorId")
		e.Log()
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}
	if len(creatorIds) < 1 {
		e := err.New("url params does not include creatorId")
		e.Log()
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	} else if len(creatorIds) > 1 {
		e := err.New("url params cannot have more than one creatorId")
		e.Log()
		http.Error(w, e.E.Error(), http.StatusBadRequest)
	}

	creator, cErr := queries.GetCreator(ctx, neo, creatorIds[0])
	if cErr.E != nil {
		cErr.Log()
		http.Error(w, cErr.E.Error(), http.StatusInternalServerError)
		return
	}

	var returnCreator ResponseCreator
	if e := utils.StructToStruct(creator, &returnCreator); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	vErr := returnCreator.validateResponseCreator()
	if vErr.E != nil {
		vErr.Log()
		http.Error(w, vErr.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(returnCreator); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, newE.E.Error(), http.StatusInternalServerError)
		return
	}
}

// GET USER CREATORS
func GetUserCreators(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getUserCreators(w, r, ctx, neo, store)
	}
}

func getUserCreators(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	user, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	retrievedUserCreators, qErr := queries.GetUserCreators(ctx, neo, user.Uid)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), http.StatusInternalServerError)
		return
	}

	responseCreators := make([]ResponseCreator, 0)
	for _, creator := range retrievedUserCreators {
		var responseCreator ResponseCreator
		if e := utils.StructToStruct(creator, &responseCreator); e.E != nil {
			e.Log()
			http.Error(w, e.E.Error(), http.StatusInternalServerError)
			return
		}
		if e := responseCreator.validateResponseCreator(); e.E != nil {
			e.Log()
			http.Error(w, e.E.Error(), http.StatusInternalServerError)
			return
		}
		responseCreators = append(responseCreators, responseCreator)
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(responseCreators); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}

// CREATE CREATOR
func CreateCreator(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createCreator(w, r, ctx, neo, store)
	}
}

type PostedCreator struct {
	Name      string `json:"name"`
	UniqueName string `json:"uniqueName"`
	About     string `json:"about"`
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
	if e := utils.StructToStruct(postedCreator, &newCreator); e.E != nil {
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

	createdCreator, cErr := queries.CreateCreator(ctx, neo, user, creatorModel)
	if cErr.E != nil {
		cErr.Log()
		http.Error(w, cErr.E.Error(), http.StatusInternalServerError)
		return
	}

	var responseCreator ResponseCreator
	if e := utils.StructToStruct(createdCreator, &responseCreator); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	vErr := responseCreator.validateResponseCreator()
	if vErr.E != nil {
		vErr.Log()
		http.Error(w, vErr.E.Error(), http.StatusInternalServerError)
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

// UPDATE CREATOR INFO
func UpdateCreator(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateCreator(w, r, ctx, neo, store)
	}
}

type PatchedUpdatedCreatorInfo struct {
	Uid       string `json:"uid"`
	Name      string `json:"name"`
	UniqueName string `json:"uniqueName"`
	About     string `json:"about"`
}

func updateCreator(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	user, uErr := middleware.AuthenticateUser(r, store)
	if uErr.E != nil {
		uErr.Log()
		http.Error(w, uErr.E.Error(), http.StatusUnauthorized)
		return
	}

	var patchedInfo PatchedUpdatedCreatorInfo
	if e := json.NewDecoder(r.Body).Decode(&patchedInfo); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, newE.E.Error(), http.StatusInternalServerError)
		return
	}

	var incomingInfo models.IncomingUpdatedCreatorInfo
	if e := utils.StructToStruct(patchedInfo, &incomingInfo); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	updatedCreatorInfo, cErr := models.MakeUpdatedCreatorInfo(incomingInfo)
	if cErr.E != nil {
		cErr.Log()
		http.Error(w, cErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	updatedCreator, exists, status, qErr := queries.UpdateCreatorInfo(ctx, neo, updatedCreatorInfo, user.Uid)
	if qErr.E != nil && exists {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
		return
	} else if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
		return
	}

	var responseCreator ResponseCreator
	if e := utils.StructToStruct(updatedCreator, &responseCreator); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	vErr := responseCreator.validateResponseCreator()
	if vErr.E != nil {
		vErr.Log()
		http.Error(w, vErr.E.Error(), http.StatusInternalServerError)
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

// DELETE CREATOR
func DeleteCreator(w http.ResponseWriter, r *http.Request) {

}
