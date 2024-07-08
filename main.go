package main

import (
	"context"
	"fmt"
	"freecreate/internal/api/routes"
	"freecreate/internal/config"
	"os"

	"github.com/joho/godotenv"
)

func run(ctx context.Context) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	fmt.Println(os.Getenv("NEO_USER"))
	neo, neoErr := config.InitNeo(ctx)
	if neoErr != nil {
		return neoErr
	}

	mongo := config.InitMongo(ctx)
	redis := config.InitRedis()
	if err := routes.CreateRoutes(ctx,  mongo, neo, redis); err != nil {
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
