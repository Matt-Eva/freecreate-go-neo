package config

import (
	"context"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func InitNeo(ctx context.Context) (neo4j.DriverWithContext, error) {
	uri := os.Getenv("NEO_URI")
	pwd := os.Getenv("NEO_PASSWORD")
	user := os.Getenv("NEO_USER")

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, pwd, ""))
	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
