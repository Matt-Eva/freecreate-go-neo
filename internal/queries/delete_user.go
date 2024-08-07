package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func DeleteUser(ctx context.Context, neo neo4j.DriverWithContext, userId string) err.Error{
	params := map[string]any{
		"userId": userId,
	}

	query, qErr := buildDeleteUserQuery()
	if qErr.E != nil {
		return qErr
	}

	db := os.Getenv("NEO_DB")
	if db == ""{
		return err.New("could not get db env variable")
	}

	_, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return err.NewFromErr(nErr)
	}

	return err.Error{}
}

func buildDeleteUserQuery()(string, err.Error){
	userLabel, lErr := GetNodeLabel("User")
	if lErr.E != nil {
		return "", lErr
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId})", userLabel)

	deleteQuery := "DETACH DELETE u"

	query := matchQuery + deleteQuery

	return query , err.Error{}
}