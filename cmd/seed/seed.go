package main

import (
	"context"
	"fmt"
	"freecreate/internal/config"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func configureNeo(ctx context.Context)(neo4j.DriverWithContext, error){

	neo, err := config.InitNeo(ctx)
	if err != nil {
		return nil, err
	}

	err = neo.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}

	return neo, nil
}

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ctx := context.Background()

	neo, err := configureNeo(ctx)
	if err != nil{
		defer neo.Close(ctx)
		fmt.Println(err.Error())
		return
	}
	defer neo.Close(ctx)

}