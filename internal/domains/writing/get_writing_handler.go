package writing

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetWritingHandler(ctx context.Context, neo neo4j.DriverWithContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getWritingHandler(w, r, ctx, neo)
	}
}

func getWritingHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext) {
	urlParams := r.URL.Query()

	writingIds, ok := urlParams["writingId"]
	if !ok {
		e := err.New("url params does not include writing id")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}

	if len(writingIds) != 1 {
		e := err.New("invalid number of writing id params - must be 1 and only 1")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}

	writingId := writingIds[0]

	returnedWriting, status, qErr := GetWritingQuery(ctx, neo, writingId)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
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
