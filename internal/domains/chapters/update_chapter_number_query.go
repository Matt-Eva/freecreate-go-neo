package chapters

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func UpdateChapterNumberQuery(ctx context.Context, mongo *mongo.Client, userId, writingId, chapterId string, newNumber int)(int, int, err.Error){
	db := os.Getenv("MONGO_DB")
	if db == ""{
		return 0, 500, err.New("db env variable is empty")
	}

	fmt.Println("userId", userId)
	fmt.Println("writingId", writingId)
	fmt.Println("chapterId", chapterId)
	fmt.Println("newNumber", newNumber)

	coll := mongo.Database(db).Collection("chapters")
	filter := bson.M{"writing_id": writingId, "user_id": userId, "uid": chapterId}
	update := bson.M{"$set": bson.M{"chapter_number": newNumber}}

	result, mErr := coll.UpdateOne(ctx, filter, update)
	if mErr != nil {
		n := err.NewFromErr(mErr)
		return 0, 500, n
	}
	if result.MatchedCount == 0 {
		return 0, 404, err.New("record not found")
	}
	if result.ModifiedCount == 0 {
		return 0, 422, err.New("record not updated")
	}

	fmt.Println(result)

	return newNumber, 200, err.Error{}
}