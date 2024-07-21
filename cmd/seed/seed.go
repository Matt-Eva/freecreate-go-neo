package main

import (
	"context"
	"freecreate/internal/config"
	"freecreate/internal/err"
	"freecreate/internal/seeds"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	lErr := godotenv.Load()
	if lErr != nil {
		e := err.NewFromErr(lErr)
		e.Log()
		os.Exit(1)
	}

	arg := os.Args[0]

	ctx := context.Background()

	neo, iErr := config.InitNeo(ctx)
	if iErr.E != nil {
		defer neo.Close(ctx)
		iErr.Log()
		os.Exit(1)
	}
	defer neo.Close(ctx)

	mongo, mErr := config.InitMongo(ctx)
	if mErr.E != nil {
		defer config.MongoDisconnect(mongo, ctx)
		mErr.Log()
		os.Exit(1)
	}
	defer config.MongoDisconnect(mongo, ctx)

	dErr := seeds.DeleteSeeds(ctx, neo)
	if dErr.E != nil {
		dErr.Log()
		os.Exit(1)
	}

	uErr := seeds.SeedUsers(neo, ctx)
	if uErr.E != nil {
		uErr.Log()
		os.Exit(1)
	}

	cErr := seeds.SeedCreators(ctx, neo)
	if cErr.E != nil {
		cErr.Log()
		os.Exit(1)
	}

}
