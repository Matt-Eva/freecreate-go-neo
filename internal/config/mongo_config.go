package config

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(ctx context.Context) (*mongo.Client, err.Error) {
	fmt.Println("connecting mongo")
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_TOKEN")).SetServerAPIOptions(serverApi)

	client, cErr := mongo.Connect(ctx, opts)
	if cErr != nil {
		e := err.NewFromErr(cErr)
		return client, e
	}

	pErr := client.Database("freecreate").RunCommand(ctx, bson.D{{"ping", 1}}).Err()
	if pErr != nil {
		e := err.NewFromErr(pErr)
		return client, e
	}

	fmt.Println("mongo connected")
	return client, err.Error{}
}

func MongoDisconnect(mon *mongo.Client, ctx context.Context) {
	if mErr := mon.Disconnect(ctx); mErr != nil {
		panic(mErr)
	  }
}
