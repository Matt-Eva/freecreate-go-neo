package main

import (
	"context"
	"fmt"
	"freecreate/internal/api/routes"
	"freecreate/internal/config"
	"os"
)

func run(ctx context.Context) error {
	neo := config.InitNeo(ctx)
	mongo := config.InitMongo(ctx)
	redis := config.InitRedis(ctx)
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
