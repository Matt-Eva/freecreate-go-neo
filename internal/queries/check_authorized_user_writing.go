package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func checkAuthorizedUserWriting(ctx context.Context, neo neo4j.DriverWithContext, userId, writingId string) (int, err.Error) {
	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return 500, uErr
	}

	isCreatorLabel, iErr := GetRelationshipLabel("IS_CREATOR")
	if iErr.E != nil {
		return 500, iErr
	}

	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return 500, cErr
	}

	createdLabel, crErr := GetRelationshipLabel("CREATED")
	if crErr.E != nil {
		return 500, crErr
	}

	writingLabel, wErr := GetNodeLabel("Writing")
	if wErr.E != nil {
		return 500, wErr
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId}), (w:%s {uid: $writingId})", userLabel, writingLabel)
	returnQuery := fmt.Sprintf("RETURN exists((u) - [:%s] -> (:%s) - [:%s] -> (w)) AS exists", isCreatorLabel, creatorLabel, createdLabel)
	query := matchQuery + returnQuery

	params := map[string]any{
		"userId":    userId,
		"writingId": writingId,
	}

	db := os.Getenv("NEO_DB")
	if db == "" {
		return 500, err.New("NEO_DB environment variable returned empty string")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return 500, err.NewFromErr(nErr)
	}
	value, ok := result.Records[0].Get("exists")
	if !ok {
		return 500, err.New("returned record does not have exists attribute")
	}

	exists, ok := value.(bool)
	if !ok {
		return 500, err.New("exists value from database could not be converted to boolean")
	}

	if exists {
		return 200, err.Error{}
	} else {
		return 401, err.New("user does not have access to this creator profile")
	}
}
