package queries

import (
	"context"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)

type RetrievedChapterNoContent struct {
	Uid string
	Title string
	WritingId string
	ChapterNumber int
}

func CreateChapter(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, chapter models.Chapter, userId string)(RetrievedChapterNoContent, int, err.Error){
	status, uErr := checkAuthorizedUserWriting(ctx, neo, userId, chapter.WritingId)
	if uErr.E != nil {
		return RetrievedChapterNoContent{}, status, uErr
	}

	db := os.Getenv("MONGO_DB")
	if db == ""{
		return RetrievedChapterNoContent{}, 500, err.New("db environment variable empty")
	}

	coll := mongo.Database(db).Collection("chapters")

	_, mErr := coll.InsertOne(ctx, chapter)
	if mErr != nil {
		return RetrievedChapterNoContent{}, 500, err.NewFromErr(mErr)
	}

	var createdChapter RetrievedChapterNoContent
	if e:= utils.StructToStruct(chapter, &createdChapter); e.E != nil {
		return createdChapter, 500, e
	}

	return createdChapter, 200, err.Error{}

}

