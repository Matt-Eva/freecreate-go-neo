package writing_handler

import (
	"context"
	"encoding/json"
	"freecreate/internal/api/middleware"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"freecreate/internal/utils"
	"net/http"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

type PostedWriting struct {
	CreatorId   string   `json:"creatorId`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	WritingType string   `json:"writingType"`
	Font        string   `json:"font"`
	Genres      []string `json:"genres'`
	Tags        []string `json:"tags"`
}

func CreateWriting(ctx context.Context, neo neo4j.DriverWithContext,  store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleCreateWriting(w, r, ctx, neo,  store)
	}
}

func handleCreateWriting(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext,  store *redisstore.RedisStore) {
	user, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	var postedWriting PostedWriting
	if e := json.NewDecoder(r.Body).Decode(&postedWriting); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, newE.E.Error(), http.StatusInternalServerError)
		return
	}

	var postedWritingModel models.PostedWriting
	if e := utils.StructToStruct(postedWriting, &postedWritingModel); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	year := time.Now().Year()
	writingModel, gErr := models.MakeWriting(postedWritingModel, year)
	if gErr.E != nil {
		gErr.Log()
		http.Error(w, gErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	createdWriting, status, wErr := queries.CreateWriting(ctx, neo, user.Uid, writingModel, postedWriting.Genres, postedWriting.Tags)
	if wErr.E != nil {
		wErr.Log()
		http.Error(w, wErr.E.Error(), status)
		return
	}

	returnedWriting := &ReturnedWriting{}

	returnedWriting.Author = createdWriting.Author
	returnedWriting.CreatorId = createdWriting.CreatorId
	returnedWriting.Title = createdWriting.Title
	returnedWriting.Description = createdWriting.Description
	returnedWriting.UniqueAuthorName = createdWriting.UniqueAuthorName
	returnedWriting.Font = createdWriting.Font
	returnedWriting.Genres = createdWriting.Genres
	returnedWriting.Tags = createdWriting.Tags
	returnedWriting.Uid = createdWriting.Uid
	returnedWriting.WritingType = createdWriting.WritingType

	if e := validateReturnedWriting((*returnedWriting), postedWriting.Genres, postedWriting.Tags); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(returnedWriting); e != nil {
		newFromE := err.NewFromErr(e)
		newFromE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}

}
