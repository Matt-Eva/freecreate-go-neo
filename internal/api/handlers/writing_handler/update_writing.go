package writing_handler

import (
	"context"
	"encoding/json"
	"freecreate/internal/api/middleware"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

type patchedWriting struct {
	Uid string	`json:"uid"`
	CreatorId string `json:"creatorId"`
	Title string `json:"title"`
	Description string `json:"description"`
	Genres []string `json:"genres"`
	Tags []string `json:"tags"`
	Font string `json:"font"`
	WritingType string `json:"writingType"`
}

func UpdateWriting(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateWriting(w, r, ctx, neo, store)
	}
}

func updateWriting(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	user, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	var patchedWriting patchedWriting
	if e := json.NewDecoder(r.Body).Decode(&patchedWriting); e != nil {
		nE := err.NewFromErr(e)
		nE.Log()
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	patchedWritingModel :=  &models.PatchedWriting{}
	patchedWritingModel.Uid = patchedWriting.Uid
	patchedWritingModel.CreatorId = patchedWriting.CreatorId
	patchedWritingModel.Title = patchedWriting.Title
	patchedWritingModel.Description = patchedWriting.Description
	patchedWritingModel.Genres = patchedWriting.Genres
	patchedWritingModel.Tags = patchedWriting.Tags
	patchedWritingModel.Font = patchedWriting.Font

	updateWritingModel, mErr := models.MakeUpdateWriting(*patchedWritingModel)
	if mErr.E != nil {
		mErr.Log()
		http.Error(w, mErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	updatedWriting, status, qErr := queries.UpdateWriting(ctx, neo, user.Uid, updateWritingModel)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
		return
	}

	returnedWriting, rErr := convertRetrievwedWritingToReturnedWriting(updatedWriting)
	if rErr.E != nil {
		rErr.Log()
		http.Error(w, rErr.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e:= json.NewEncoder(w).Encode(returnedWriting); e!= nil {
		nE := err.NewFromErr(e)
		nE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
