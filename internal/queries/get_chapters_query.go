package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func GetChapters(ctx context.Context, mongo *mongo.Client, writingId string)([]RetrievedChapterNoContent, err.Error){
	db := os.Getenv("MONGO_DB")
	if db == ""{
		return []RetrievedChapterNoContent{}, err.New("db environment variable empty")
	}

	coll := mongo.Database(db).Collection("chapters")
	filter := bson.M{"writing_id": writingId}
	sort := bson.M{"chapter_number": 1}
	opts := options.Find().SetSort(sort)

	cursor, mErr := coll.Find(ctx, filter, opts)
	if mErr != nil {
		return []RetrievedChapterNoContent{}, err.NewFromErr(mErr)
	}

	results := []RetrievedChapterNoContent{}
	if e := cursor.All(ctx, &results); e != nil {
		return []RetrievedChapterNoContent{}, err.NewFromErr(e)
	}



	fmt.Println("results", results)

	return results, err.Error{}
}