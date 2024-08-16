package writing_handler

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"freecreate/internal/queries"
	"freecreate/internal/utils"
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

	retrievedWriting, qErr := queries.GetWriting(ctx, neo, mongo, creatorId, writingId)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), http.StatusInternalServerError)
		return
	}

	var returnedWriting ReturnedWriting
	if e := utils.StructToStruct(retrievedWriting, &returnedWriting); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(returnedWriting); e != nil {
		newE := err.NewFromErr(e)
		newE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}
