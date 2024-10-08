package chapters

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)


func GetChaptersHandler(ctx context.Context, mongo *mongo.Client) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		getChaptersHandler(w, r, ctx, mongo)
	}
}

func getChaptersHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, mongo *mongo.Client){
	params := r.URL.Query()
	idSlice, ok := params["writingId"]
	if !ok {
		e := err.New("params do not include writing id field")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}
	if len(idSlice) < 1 {
		e := err.New("no writing id passed in url query params")
		http.Error(w, e.E.Error(), http.StatusBadRequest)
		return
	}

	writingId := idSlice[0]

	results, mErr := GetChapters(ctx, mongo, writingId)
	if mErr.E != nil {
		mErr.Log()
		http.Error(w, mErr.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(results); e != nil {
		ne := err.NewFromErr(e)
		ne.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}