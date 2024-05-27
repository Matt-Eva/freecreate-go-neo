package main

import (
	"context"
	"fmt"
	"freecreate/internal/api/routes"
	"freecreate/internal/database"
	"os"
)

func run(ctx context.Context) error {
	neo := database.InitNeo(ctx)
	mongo := database.InitMongo(ctx)
	if err := routes.CreateRoutes(neo, mongo); err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
