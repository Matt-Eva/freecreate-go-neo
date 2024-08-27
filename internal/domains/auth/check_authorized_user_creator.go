package auth

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CheckAuthorizedUserCreator(ctx context.Context, neo neo4j.DriverWithContext, userId, creatorId string) (int, err.Error) {
	userLabel, uErr := utils.GetNodeLabel("User")
	if uErr.E != nil {
		return 500, uErr
	}

	isCreatorLabel, iErr := utils.GetRelationshipLabel("IS_CREATOR")
	if iErr.E != nil {
		return 500, iErr
	}

	creatorLabel, cErr := utils.GetNodeLabel("Creator")
	if cErr.E != nil {
		return 500, cErr
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId}), (c:%s {uid: $creatorId})", userLabel, creatorLabel)
	returnQuery := fmt.Sprintf("RETURN exists((u) - [:%s] -> (c)) AS exists", isCreatorLabel)
	query := matchQuery + returnQuery

	params := map[string]any{
		"userId":    userId,
		"creatorId": creatorId,
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
