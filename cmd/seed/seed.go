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

	uErr := seeds.SeedUsers(neo, ctx)
	if uErr !=nil {
		fmt.Fprintf(os.Stderr, "%s\n", uErr)
		os.Exit(1)
	}

}
