package writing_handler

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"freecreate/internal/queries"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetWriting(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getWriting(w, r, ctx, neo, mongo)
	}
}

func getWriting(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client) {
	urlParams := r.URL.Query()
	creatorIds, ok := urlParams["creatorId"]
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
	if len(writingIds) != 1 {
		e := err.New("invalid number of writing id params - must be 1 and only 1")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}

	creatorId := creatorIds[0]
	writingId := writingIds[0]

	retrievedWriting, status, qErr := queries.GetWriting(ctx, neo, creatorId, writingId)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
		return
	}

	returnedWriting := &ReturnedWriting{}

	returnedWriting.Author = retrievedWriting.Author
	returnedWriting.Uid = retrievedWriting.Uid
	returnedWriting.Description = retrievedWriting.Description
	returnedWriting.Title = retrievedWriting.Title
	returnedWriting.Genres = retrievedWriting.Genres
	returnedWriting.Tags = retrievedWriting.Tags
	returnedWriting.CreatorId = retrievedWriting.CreatorId
	returnedWriting.UniqueAuthorName = retrievedWriting.UniqueAuthorName
	returnedWriting.Font = retrievedWriting.Font

	if e := validateReturnedWriting(*returnedWriting, retrievedWriting.Genres, retrievedWriting.Tags); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusUnprocessableEntity)
		return
	}
	

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode((*returnedWriting)); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}
