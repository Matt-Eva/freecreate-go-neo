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

func UpdateCreatorInfo(ctx context.Context, neo neo4j.DriverWithContext, info models.UpdatedCreatorInfo) (UpdatedCreator, bool, err.Error) {
	exists, uErr := checkUniqueCreatorId(ctx, neo, info.CreatorId)
	if uErr.E != nil && exists {
		return UpdatedCreator{}, true, uErr
	} else if uErr.E != nil {
		return UpdatedCreator{}, false, uErr
	}

	params := buildUpdateCreatorInfoParams(info)
	query, qErr := buildUpdateCreatorInfoQuery(params)
	if qErr.E != nil {
		return UpdatedCreator{}, false, qErr
	}

	db := os.Getenv("NEO_DB")
	if db == "" {
		return UpdatedCreator{}, false, err.New("NEO_DB environment variable returned empty string")
	}
	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return UpdatedCreator{},false, e
	}
	if len(result.Records) < 1 {
		return UpdatedCreator{},false, err.New("db query returned zero records")
	}

	resultMap := result.Records[0].AsMap()
	var updatedCreator UpdatedCreator
	if e := utils.MapToStruct(resultMap, &updatedCreator); e.E != nil {
		return UpdatedCreator{}, false, e
	}

	return updatedCreator, false, err.Error{}
}

func checkUniqueCreatorId(ctx context.Context, neo neo4j.DriverWithContext, creatorId string)(bool, err.Error){
	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return false, cErr
	}

	query := fmt.Sprintf("MATCH (c:%s {creatorId: $creatorId}) RETURN c.creatorId AS CreatorId", creatorLabel)
	params := map[string]any {
		"creatorId": creatorId,
	}

	db := os.Getenv("NEO_DB")
	if db == ""{
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

	matchQuery := fmt.Sprintf("MATCH (c:%s {uid: $uid})", creatorLabel)

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
		if i < len(setAttributes) - 1{
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
