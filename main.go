package main

import (
	"context"
	"fmt"
	"freecreate/internal/api/routes"
	"freecreate/internal/databases"
	"os"
)

func run(ctx context.Context) error {
	neo := databases.InitNeo(ctx)
	mongo := databases.InitMongo(ctx)
	redis := databases.InitRedis(ctx)
	if err := routes.CreateRoutes(neo, mongo, redis); err != nil {
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
