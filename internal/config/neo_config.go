package config

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func InitNeo(ctx context.Context) (neo4j.DriverWithContext, err.Error) {
	fmt.Println("connecting neo")
	uri := os.Getenv("NEO_URI")
	pwd := os.Getenv("NEO_PASSWORD")
	user := os.Getenv("NEO_USER")

	driver, nErr := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, pwd, ""))
	if nErr != nil {
		defer driver.Close(ctx)
		e := err.NewFromErr(nErr)
		return nil, e
	}

	cErr := driver.VerifyConnectivity(ctx)
	if cErr != nil {
		defer driver.Close(ctx)
		e := err.NewFromErr(cErr)
		return nil, e
	}

	fmt.Println("neo connected")
	return driver, err.Error{}
}
