package config

import (
	"context"
	"freecreate/internal/err"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(ctx context.Context) (*mongo.Client, err.Error) {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_TOKEN")).SetServerAPIOptions(serverApi)

	client, cErr := mongo.Connect(ctx, opts)
	if cErr != nil {
		e := err.NewFromErr(cErr)
		return client, e
	}

	return client, err.Error{}
}
