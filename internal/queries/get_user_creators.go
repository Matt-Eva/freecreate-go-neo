package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type RetrievedUserCreator struct {
	Name      string
	Uid       string
	UniqueName string
	About     string
}

func (r RetrievedUserCreator) validatedRetrievedUserCreator() err.Error {
	if r.Name == "" {
		return err.New("retrieved user creator name cannot be empty")
	}
	if r.Uid == "" {
		return err.New("retrieved user creator Uid cannot be empty")
	}
	if r.UniqueName == "" {
		return err.New("retrieved user creator UniqueName cannot be empty")
	}
	return err.Error{}
}

func GetUserCreators(ctx context.Context, neo neo4j.DriverWithContext, userId string) ([]RetrievedUserCreator, err.Error) {
	query, qErr := buildGetUserCreatorsQuery()
	if qErr.E != nil {
		return []RetrievedUserCreator{}, qErr
	}
	params := map[string]any{
		"userId": userId,
	}

	db := os.Getenv("NEO_DB")
	if db == "" {
		return []RetrievedUserCreator{}, err.New("neo db environment variable is empty")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return []RetrievedUserCreator{}, err.NewFromErr(nErr)
	}

	var retrievedCreators []RetrievedUserCreator
	for _, record := range result.Records {
		recordMap := record.AsMap()
		var retrievedCreator RetrievedUserCreator
		if e := utils.MapToStruct(recordMap, &retrievedCreator); e.E != nil {
			return []RetrievedUserCreator{}, e
		}
		vErr := retrievedCreator.validatedRetrievedUserCreator()
		if vErr.E != nil {
			return []RetrievedUserCreator{}, vErr
		}
		retrievedCreators = append(retrievedCreators, retrievedCreator)
	}

	return retrievedCreators, err.Error{}
}

func buildGetUserCreatorsQuery() (string, err.Error) {
	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	isCreatorLabel, iErr := GetRelationshipLabel("IS_CREATOR")
	if iErr.E != nil {
		return "", iErr
	}

	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId}) -[:%s] -(c:%s)", userLabel, isCreatorLabel, creatorLabel)
	returnQuery := `RETURN c.name AS Name, c.uid AS Uid, c.about AS About, c.uniqueName AS UniqueName`
	query := matchQuery + returnQuery

	return query, err.Error{}
}
