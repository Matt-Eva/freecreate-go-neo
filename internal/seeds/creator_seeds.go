package seeds

import (
	"context"
	"fmt"
	"freecreate/internal/err"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SeedCreators(ctx context.Context, neo neo4j.DriverWithContext) err.Error {
	result, uErr := getSeedUsers(ctx, neo)
	if uErr.E != nil {
		return uErr
	}

	for _, user := range result {
		uid := user["uid"]
		fmt.Println(uid)
	}

	return err.Error{}
}

func getSeedUsers(ctx context.Context, neo neo4j.DriverWithContext) ([]map[string]any, err.Error) {
	users := []map[string]any{}

	query := `
		MATCH (u:User)
		WHERE u.seed = true
		RETURN u.uid AS uidd
	`

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return []map[string]any{}, e
	}
	if len(result.Records) < 1 {
		e := err.New("no seed users in database")
		return []map[string]any{}, e
	}

	for _, record := range result.Records {
		rMap := record.AsMap()
		users = append(users, rMap)
	}

	return users, err.Error{}
}

func seedCreator(ctx context.Context, neo neo4j.DriverWithContext) {

}

func makeSeedCreator() {

}
