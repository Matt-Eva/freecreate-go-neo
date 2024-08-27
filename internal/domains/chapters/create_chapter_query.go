package chapters

import (
	"context"
	"freecreate/internal/domains/auth"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)



func CreateChapter(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, chapter Chapter, userId string) (ReturnChapterNoContent, int, err.Error) {
	status, uErr := auth.CheckAuthorizedUserWriting(ctx, neo, userId, chapter.WritingId)
	if uErr.E != nil {
		return ReturnChapterNoContent{}, status, uErr
	}

	db := os.Getenv("MONGO_DB")
	if db == "" {
		return ReturnChapterNoContent{}, 500, err.New("db environment variable empty")
	}

	coll := mongo.Database(db).Collection("chapters")

	_, mErr := coll.InsertOne(ctx, chapter)
	if mErr != nil {
		return ReturnChapterNoContent{}, 500, err.NewFromErr(mErr)
	}

	var createdChapter ReturnChapterNoContent
	if e := utils.StructToStruct(chapter, &createdChapter); e.E != nil {
		return createdChapter, 500, e
	}

	return createdChapter, 200, err.Error{}

}
