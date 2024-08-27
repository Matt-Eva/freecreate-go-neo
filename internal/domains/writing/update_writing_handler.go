package writing

import (
	"context"
	"encoding/json"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/middleware"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

type patchedWriting struct {
	Uid         string   `json:"uid"`
	CreatorId   string   `json:"creatorId"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
	Tags        []string `json:"tags"`
	Font        string   `json:"font"`
	WritingType string   `json:"writingType"`
}

func UpdateWritingHandler(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		updateWritingHandler(w, r, ctx, neo, store)
	}
}

func updateWritingHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
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

	fmt.Println("patchedWriting", patchedWriting)

	patchedWritingModel := &PatchedWriting{}
	patchedWritingModel.Uid = patchedWriting.Uid
	patchedWritingModel.CreatorId = patchedWriting.CreatorId
	patchedWritingModel.Title = patchedWriting.Title
	patchedWritingModel.Description = patchedWriting.Description
	patchedWritingModel.Genres = patchedWriting.Genres
	patchedWritingModel.Tags = patchedWriting.Tags
	patchedWritingModel.Font = patchedWriting.Font
	patchedWritingModel.WritingType = patchedWriting.WritingType

	updateWritingModel, mErr := MakeUpdateWriting(*patchedWritingModel)
	if mErr.E != nil {
		mErr.Log()
		http.Error(w, mErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	updatedWriting, status, qErr := UpdateWritingQuery(ctx, neo, user.Uid, updateWritingModel)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(updatedWriting); e != nil {
		nE := err.NewFromErr(e)
		nE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
