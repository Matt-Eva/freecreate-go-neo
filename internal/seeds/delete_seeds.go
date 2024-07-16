package seeds

import (
	"context"
	"errors"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func DeleteSeeds(ctx context.Context, neo neo4j.DriverWithContext) error {
	deleteQuery := `
		MATCH(n)
		WHERE n.seed = true
		DETACH DELETE n
		RETURN n.seed AS seed, labels(n) AS labels
	`

	result, err := neo4j.ExecuteQuery(ctx, neo, deleteQuery, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return err
	}

	if len(result.Records) < 1 {
		fmt.Println("seed data already empty - no nodes deleted")
	} else {
		for _, record := range result.Records {
			isSeed, ok := record.Get("seed")
			if !ok {
				return errors.New("deleted seed record did not have 'seed' field upon return")
			}

			labels, ok := record.Get("labels")
			if !ok {
				return errors.New("delete seed record did not have 'labels' field upon return")
			}

			fmt.Println("deleted record isSeed", isSeed)
			fmt.Println("deleted record labels", labels)
		}
	}

	return nil

}