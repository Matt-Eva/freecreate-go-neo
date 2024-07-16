package seeds

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SeedCreators(ctx context.Context, neo neo4j.DriverWithContext) error{
	result, uErr := getUsers(ctx, neo)
	if uErr != nil {
		return uErr
	}

	for _, user := range result {
		uid := user["uid"]
		fmt.Println(uid)
	}

	return nil
}

func getUsers(ctx context.Context, neo neo4j.DriverWithContext)([]map[string]string, error){
	users := []map[string]string{}

	// query := `
	// 	MATCH (u:User)
	// 	RETURN u.uid AS uidd
	// `

	// result, nErr :

	return users, nil
}

func seedCreator(ctx context.Context, neo neo4j.DriverWithContext) {

}

func makeSeedCreator() {

}
