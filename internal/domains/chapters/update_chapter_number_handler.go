package chapters

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"freecreate/internal/middleware"
	"net/http"

	"github.com/rbcervilla/redisstore/v9"
	"go.mongodb.org/mongo-driver/mongo"
)


func UpdateChapterNumberHandler(ctx context.Context, mongo *mongo.Client, store *redisstore.RedisStore)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		handleUpdateChapterNumber(w, r, ctx, mongo, store)
	}
}

type PatchedChapterNumber struct {
	NewNumber int `json:"newNumber"`
	ChapterId string `json:"chapterId"`
	WritingId string `json:"writingId"`
}

func handleUpdateChapterNumber(w http.ResponseWriter, r *http.Request, ctx context.Context, mongo *mongo.Client, store *redisstore.RedisStore){
	user, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	var patchedChapterNumber PatchedChapterNumber
	if e := json.NewDecoder(r.Body).Decode(&patchedChapterNumber); e !=nil {
		n := err.NewFromErr(e)
		n.Log()
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	newNumber, status, qErr := UpdateChapterNumberQuery(ctx, mongo, user.Uid, patchedChapterNumber.WritingId, patchedChapterNumber.ChapterId, patchedChapterNumber.NewNumber)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
	}

	type updateChapterReturnStruct struct{
		NewNumber int `json:"newNumber"`
	}

	updatedNumber := updateChapterReturnStruct{
		NewNumber: newNumber,
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(updatedNumber); e != nil {
		n := err.NewFromErr(e)
		n.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}