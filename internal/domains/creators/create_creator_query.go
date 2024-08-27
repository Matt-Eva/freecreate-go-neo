package creators

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/middleware"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type CreatedCreator struct {
	Uid        string
	Name       string
	UniqueName string
	About      string
}

func CreateCreatorQuery(ctx context.Context, neo neo4j.DriverWithContext, user middleware.AuthenticatedUser, creator Creator) (CreatedCreator, err.Error) {
	uErr := checkUniqueCreator(ctx, neo, creator)
	if uErr.E != nil {
		return CreatedCreator{}, uErr
	}

	query, qErr := buildCreateCreatorQuery()
	if qErr.E != nil {
		return CreatedCreator{}, qErr
	}

	params := buildCreateCreatorParams(user, creator)

	db := os.Getenv("NEO_DB")
	if db == "" {
		return CreatedCreator{}, err.New("NEO_DB environment variable return empty string")
	}
	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return CreatedCreator{}, err.NewFromErr(nErr)
	}
	if len(result.Records) < 1 {
		return CreatedCreator{}, err.New("Create creator query returned zero records")
	}

	resultMap := result.Records[0].AsMap()
	createdCreator := CreatedCreator{}
	cErr := utils.MapToStruct(resultMap, &createdCreator)
	if cErr.E != nil {
		return CreatedCreator{}, cErr
	}

	return createdCreator, err.Error{}
}

func checkUniqueCreator(ctx context.Context, neo neo4j.DriverWithContext, creator Creator) err.Error {
	query, qErr := buildCheckUniqueCreatorQuery()
	if qErr.E != nil {
		return qErr
	}

	params := buildCheckUniqueCreatorParams(creator)

	db := os.Getenv("NEO_DB")
	if db == "" {
		return err.New("NEO_DB environment variable return empty string")
	}
	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return err.NewFromErr(nErr)
	}

	if len(result.Records) > 0 {
		return err.New("creatorId already exists")
	}

	return err.Error{}
}

func buildCheckUniqueCreatorQuery() (string, err.Error) {
	creatorLabel, cErr := utils.GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	matchQuery := fmt.Sprintf("MATCH (c:%s {uniqueName: $uniqueName})", creatorLabel)
	returnQuery := `RETURN c.uid AS Uid`
	query := matchQuery + returnQuery

	return query, err.Error{}
}

func buildCheckUniqueCreatorParams(creator Creator) map[string]any {
	return map[string]any{
		"uniqueName": creator.UniqueName,
	}
}

func buildCreateCreatorQuery() (string, err.Error) {
	creatorLabel, cErr := utils.GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	userLabel, uErr := utils.GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	isCreatorLabel, rErr := utils.GetRelationshipLabel("IS_CREATOR")
	if rErr.E != nil {
		return "", rErr
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId})", userLabel)
	createQuery := fmt.Sprintf("CREATE (c:%s $creatorParams) <-[r:%s]-(u)", creatorLabel, isCreatorLabel)
	returnQuery := `RETURN c.uid AS Uid, c.name AS Name, c.about AS About, c.uniqueName AS UniqueName`
	query := matchQuery + createQuery + returnQuery
	return query, err.Error{}
}

func buildCreateCreatorParams(user middleware.AuthenticatedUser, creator Creator) map[string]any {
	creatorParams := utils.StructToMap(creator)

	return map[string]any{
		"userId":        user.Uid,
		"creatorParams": creatorParams,
	}
}
