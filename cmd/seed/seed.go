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
	ctx := context.Background()

	neo, iErr := config.InitNeo(ctx)
	if iErr != nil {
		defer neo.Close(ctx)
		e := err.NewFromErr(iErr)
		e.Log()
		os.Exit(1)
	}
	defer neo.Close(ctx)

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

}
