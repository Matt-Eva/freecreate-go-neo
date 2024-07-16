package seeds

import (
	"context"
	"fmt"
	"freecreate/internal/err"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func DeleteSeeds(ctx context.Context, neo neo4j.DriverWithContext) err.Error {
	deleteQuery := `
		MATCH(n)
		WHERE n.seed = true
		DETACH DELETE n
		RETURN n.seed AS seed, labels(n) AS labels
	`

	result, eErr := neo4j.ExecuteQuery(ctx, neo, deleteQuery, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if eErr != nil {
		e := err.NewFromErr(eErr)
		return e
	}

	if len(result.Records) < 1 {
		fmt.Println("seed data already empty - no nodes deleted")
	} else {
		for _, record := range result.Records {
			isSeed, ok := record.Get("seed")
			if !ok {
				return err.New("deleted seed record did not have 'seed' field upon return")
			}

			labels, ok := record.Get("labels")
			if !ok {
				return err.New("delete seed record did not have 'labels' field upon return")
			}

			fmt.Println("deleted record isSeed", isSeed)
			fmt.Println("deleted record labels", labels)
		}
	}

	return err.Error{}

}
