package chapter_handler

import (
	"context"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

func CreateChapter(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
return func (w http.ResponseWriter, r *http.Request){}
}

type PostedChapter struct {
	Title string `json:"title"`
	ChapterNumber int `json:"chapterNumber"`
	WritingId string `json:"writingId"`
}

func createChapter(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore){
	
}