package seeds

import (
	"context"
	"fmt"
	"freecreate/internal/config"
)

func SeedUsers(){
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