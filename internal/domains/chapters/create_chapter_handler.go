package chapters

import (
	"context"
	"encoding/json"
	"freecreate/internal/err"
	"freecreate/internal/middleware"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"freecreate/internal/utils"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChapterHandler(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createChapterHandler(w, r, ctx, neo, mongo, store)
	}
}

type PostedChapter struct {
	Title         string `json:"title"`
	ChapterNumber int    `json:"chapterNumber"`
	WritingId     string `json:"writingId"`
}

func createChapterHandler(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, store *redisstore.RedisStore) {
	user, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		aErr.Log()
		http.Error(w, aErr.E.Error(), http.StatusUnauthorized)
		return
	}

	var postedChapter PostedChapter
	if e := json.NewDecoder(r.Body).Decode(&postedChapter); e != nil {
		newFromE := err.NewFromErr(e)
		newFromE.Log()
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	var postedChapterModel models.PostedChapter
	if e := utils.StructToStruct(postedChapter, &postedChapterModel); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	chapterModel, vErr := models.MakeChapter(postedChapterModel)
	if vErr.E != nil {
		vErr.Log()
		http.Error(w, vErr.E.Error(), http.StatusUnprocessableEntity)
		return
	}

	chapter, status, qErr := queries.CreateChapter(ctx, neo, mongo, chapterModel, user.Uid)
	if qErr.E != nil {
		qErr.Log()
		http.Error(w, qErr.E.Error(), status)
		return
	}

	var returnChapter ReturnChapterNoContent
	if e := utils.StructToStruct(chapter, &returnChapter); e.E != nil {
		e.Log()
		http.Error(w, e.E.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if e := json.NewEncoder(w).Encode(returnChapter); e != nil {
		nE := err.NewFromErr(e)
		nE.Log()
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
}
