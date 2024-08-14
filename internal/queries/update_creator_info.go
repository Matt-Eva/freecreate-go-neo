package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type UpdatedCreator struct {
	Uid       string
	Name      string
	CreatorId string
	About     string
}

// return creator, if creatorId already exists, http status, error
func UpdateCreatorInfo(ctx context.Context, neo neo4j.DriverWithContext, info models.UpdatedCreatorInfo, userId string) (UpdatedCreator, bool, int, err.Error) {
	status, aErr := checkAuthorizedUser(ctx, neo, userId, info.Uid)
	if aErr.E != nil {
		return UpdatedCreator{}, false, status, aErr
	}

	exists, uErr := checkUniqueCreatorId(ctx, neo, info.CreatorId)
	if uErr.E != nil && exists {
		return UpdatedCreator{}, true, 422, uErr
	} else if uErr.E != nil {
		return UpdatedCreator{}, false, 500, uErr
	}

	params := buildUpdateCreatorInfoParams(info)
	query, qErr := buildUpdateCreatorInfoQuery(params)
	if qErr.E != nil {
		return UpdatedCreator{}, false, 500, qErr
	}
	params["userId"] = userId

	db := os.Getenv("NEO_DB")
	if db == "" {
		return UpdatedCreator{}, false, 500, err.New("NEO_DB environment variable returned empty string")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return UpdatedCreator{}, false, 500, e
	}
	if len(result.Records) < 1 {
		return UpdatedCreator{}, false, 500, err.New("db query returned zero records")
	}

	resultMap := result.Records[0].AsMap()
	var updatedCreator UpdatedCreator
	if e := utils.MapToStruct(resultMap, &updatedCreator); e.E != nil {
		return UpdatedCreator{}, false, 500, e
	}

	return updatedCreator, false, 201, err.Error{}
}

func checkAuthorizedUser(ctx context.Context, neo neo4j.DriverWithContext, userId, creatorId string) (int, err.Error) {
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

func checkUniqueCreatorId(ctx context.Context, neo neo4j.DriverWithContext, creatorId string) (bool, err.Error) {
	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return false, cErr
	}

	query := fmt.Sprintf("MATCH (c:%s {creatorId: $creatorId}) RETURN c.creatorId AS CreatorId", creatorLabel)
	params := map[string]any{
		"creatorId": creatorId,
	}

	db := os.Getenv("NEO_DB")
	if db == "" {
		return false, err.New("neo db env variable is empty")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return false, err.NewFromErr(nErr)
	}
	if len(result.Records) > 0 {
		msg := fmt.Sprintf("creator id '%s' already in use", creatorId)
		return true, err.New(msg)
	}

	return false, err.Error{}
}

func buildUpdateCreatorInfoQuery(params map[string]any) (string, err.Error) {
	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	isCreatorLabel, iErr := GetRelationshipLabel("IS_CREATOR")
	if iErr.E != nil {
		return "", iErr
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId}) - [:%s] -> (c:%s {uid: $uid})", userLabel, isCreatorLabel, creatorLabel)

	type AttrStruct struct {
		Key       string
		Attribute string
	}

	var setAttributes []AttrStruct
	for key := range params {
		attribute := "$" + key
		attrMap := AttrStruct{
			Attribute: attribute,
			Key:       key,
		}
		setAttributes = append(setAttributes, attrMap)
	}

	setQuery := "SET "
	for i, attrMap := range setAttributes {
		query := ""
		if i < len(setAttributes)-1 {
			query = fmt.Sprintf("c.%s = %s, ", attrMap.Key, attrMap.Attribute)
		} else {
			query = fmt.Sprintf("c.%s = %s ", attrMap.Key, attrMap.Attribute)
		}
		setQuery += query
	}

	returnQuery := `
	RETURN c.uid AS Uid, c.name AS Name, c.creatorId AS CreatorId, c.about AS About
	`

	query := matchQuery + setQuery + returnQuery

	return query, err.Error{}
}

func buildUpdateCreatorInfoParams(info models.UpdatedCreatorInfo) map[string]any {
	params := utils.StructToMap(info)
	return params
}
