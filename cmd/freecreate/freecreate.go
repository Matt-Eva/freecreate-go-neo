package main

import (
	"context"
	"freecreate/internal/api/routes"
	"freecreate/internal/config"
	"freecreate/internal/err"
	"os"

	"github.com/joho/godotenv"
)

func run(ctx context.Context) err.Error {
	lErr := godotenv.Load()
	if lErr != nil {
		e := err.NewFromErr(lErr)
		return e
	}

	neo, neoErr := config.InitNeo(ctx)
	if neoErr.E != nil {
		defer neo.Close(ctx)
		return neoErr
	}

	mongo, mErr := config.InitMongo(ctx)
	if mErr.E != nil {
		defer config.MongoDisconnect(mongo, ctx)
		return mErr
	}

	redis := config.InitRedis()
	if rErr := routes.CreateRoutes(ctx, mongo, neo, redis); rErr.E != nil {
		return rErr
	}

	return err.Error{}
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err.E != nil {
		err.Log()
		os.Exit(1)
	}
}
