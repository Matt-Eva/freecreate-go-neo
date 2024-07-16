package main

import (
	"context"
	"fmt"
	"freecreate/internal/config"
	"freecreate/internal/seeds"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ctx := context.Background()

	neo, err := config.InitNeo(ctx)
	if err != nil {
		defer neo.Close(ctx)
		fmt.Println(err.Error())
		return
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
