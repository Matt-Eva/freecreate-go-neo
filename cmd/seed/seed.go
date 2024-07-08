package main

import (
	"context"
	"fmt"
	"freecreate/internal/config"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(os.Getenv("NEO_USER"))
	ctx := context.Background()
	neo, err := config.InitNeo(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer neo.Close(ctx)

	err = neo.VerifyConnectivity(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("neo connection successful")

}