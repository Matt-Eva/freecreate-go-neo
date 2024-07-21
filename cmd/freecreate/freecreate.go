package main

import (
	"context"
	"fmt"
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
	fmt.Println(os.Getenv("NEO_USER"))
	neo, neoErr := config.InitNeo(ctx)
	if neoErr.E != nil {
		return neoErr
	}

	mongo, mErr := config.InitMongo(ctx)
	if mErr.E != nil {
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
