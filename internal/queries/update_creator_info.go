package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type IncomingUpdatedCreatorInfo struct {
	Uid string
	Name string
	CreatorId string
	About string
}

type UpdatedCreator struct {
	Uid string
	Name string
	CreatorId string
	About string
}

func UpdateCreatorInfo(ctx context.Context, neo neo4j.DriverWithContext, info IncomingUpdatedCreatorInfo)(UpdatedCreator, err.Error){
	params := buildUpdateCreatorInfoParams(info)
	query, qErr := buildUpdateCreatorInfoQuery(params)
	if qErr.E != nil {
		return UpdatedCreator{}, qErr
	}

	db := os.Getenv("NEO_DB")
	if db == ""{
		return UpdatedCreator{}, err.New("NEO_DB environment variable returned empty string")
	}
	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return UpdatedCreator{}, e
	}
	if len(result.Records) < 1 {
		return UpdatedCreator{}, err.New("db query returned zero records")
	}

	resultMap := result.Records[0].AsMap()
	var updatedCreator UpdatedCreator
	if e := utils.MapToStruct(resultMap, updatedCreator); e.E != nil {
		return UpdatedCreator{}, e
	}

	return updatedCreator, err.Error{}
}

func buildUpdateCreatorInfoQuery(params map[string]any)(string, err.Error){
	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	matchQuery := fmt.Sprintf("MATCH (c:%s {uid: $uid})", creatorLabel)

	type AttrStruct struct {
		Key string
		Attribute string
	}

	var setAttributes []AttrStruct
	for key, _ := range params{
		attribute := "$" + key
		attrMap := AttrStruct{
			Attribute: attribute,
			Key: key,
		}
		setAttributes = append(setAttributes, attrMap)
	}

	setQuery := ""
	for _, attrMap := range setAttributes{
		query := fmt.Sprintf("SET c.%s = %s", attrMap.Key, attrMap.Attribute)
		setQuery += query
	}

	returnQuery := `RETURN c.uid AS Uid, c.name AS Name, c.creatorId AS CreatorId, c.about AS About`

	query := matchQuery + setQuery + returnQuery

	return query, err.Error{}
}

func buildUpdateCreatorInfoParams(info IncomingUpdatedCreatorInfo)map[string]any{
	params := utils.StructToMap(info)
	return params
}