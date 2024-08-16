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
		DETACH DELETE n
	`

	_, eErr := neo4j.ExecuteQuery(ctx, neo, deleteQuery, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if eErr != nil {
		e := err.NewFromErr(eErr)
		return e
	}

	fmt.Println("seed nodes deleted")

	return err.Error{}

}
