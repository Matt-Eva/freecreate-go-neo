package chapters

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func GetChapters(ctx context.Context, mongo *mongo.Client, writingId string)([]ReturnChapterNoContent, err.Error){
	db := os.Getenv("MONGO_DB")
	if db == ""{
		return []ReturnChapterNoContent{}, err.New("db environment variable empty")
	}

	coll := mongo.Database(db).Collection("chapters")
	filter := bson.M{"writing_id": writingId}
	sort := bson.M{"chapter_number": 1}
	opts := options.Find().SetSort(sort)

	cursor, mErr := coll.Find(ctx, filter, opts)
	if mErr != nil {
		return []ReturnChapterNoContent{}, err.NewFromErr(mErr)
	}

	results := []ReturnChapterNoContent{}
	if e := cursor.All(ctx, &results); e != nil {
		return []ReturnChapterNoContent{}, err.NewFromErr(e)
	}



	fmt.Println("results", results)

	return results, err.Error{}
}