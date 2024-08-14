package main

import (
	"context"
	"errors"
	"fmt"
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

	var args []string
	all := len(os.Args) < 2
	if !all {
		args = os.Args[1:]
	}

	if all {
		dErr := seeds.DeleteSeeds(ctx, neo)
		if dErr.E != nil {
			dErr.Log()
			os.Exit(1)
		}
	}

	calls := map[string]bool{
		"users":    false,
		"creators": false,
	}

	for _, arg := range args {
		_, ok := calls[arg]
		if ok {
			calls[arg] = true
		} else {
			eMsg := fmt.Sprintf("arg '%s' not a valid arg", arg)
			e := errors.New(eMsg)
			nE := err.NewFromErr(e)
			nE.Log()
			os.Exit(1)
		}
	}

	if calls["users"] || all {
		uErr := seeds.SeedUsers(neo, ctx)
		if uErr.E != nil {
			uErr.Log()
			os.Exit(1)
		}
	}

	if calls["creators"] || all {
		cErr := seeds.SeedCreators(ctx, neo)
		if cErr.E != nil {
			cErr.Log()
			os.Exit(1)
		}
	}

	// if calls["short_story"] || all {
	// 	sErr := seeds.SeedShortStories(ctx, neo, mongo)
	// 	if sErr.E != nil {
	// 		sErr.Log()
	// 		os.Exit(1)
	// 	}
	// }

}
