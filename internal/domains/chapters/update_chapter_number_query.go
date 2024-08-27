package chapters

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateChapterNumberQuery(ctx context.Context, mongo *mongo.Client, userId, chapterId string, newNumber int){
	
}