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

type UpdatedUser struct {
	Uid 	string
	UserId               string
	Username             string
	Email                string
	BirthDay             int
	BirthYear            int
	BirthMonth           int
	ProfilePic           string
}

func UpdateUserInfo(ctx context.Context, neo neo4j.DriverWithContext, userId string, user models.UpdatedUserInfo)(UpdatedUser, err.Error){
	params := utils.StructToMap(user)

	query, qErr := buildUpdateUserInfoQuery(params)
	if qErr.E != nil {
		return UpdatedUser{}, qErr
	}

	params["userId"] = userId

	db := os.Getenv("NEO_DB")
	if db == ""{
		return UpdatedUser{}, err.New("could not get neo db env variable")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return UpdatedUser{}, e
	}
	if len(result.Records) < 1 {
		return UpdatedUser{}, err.New("db query returned no records")
	}

	resultMap := result.Records[0].AsMap()
	var updatedUser UpdatedUser
	if e := utils.MapToStruct(resultMap, &updatedUser); e.E != nil {
		return updatedUser, e
	}

	return updatedUser, err.Error{}
}

func buildUpdateUserInfoQuery(params map[string]any)(string, err.Error){
	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId})",userLabel)


	type AttrStruct struct {
		Key string
		Attribute string
	}

	attrList := make([]AttrStruct, 0)
	for key := range params {
		attribute := "$" + key
		attrStruct := AttrStruct{
			key,
			attribute,
		}
		attrList = append(attrList, attrStruct)
	}

	setQuery := ""
	for _, attrStruct := range attrList {
		query := fmt.Sprintf("SET u.%s = %s", attrStruct.Key, attrStruct.Attribute)
		setQuery += query
	}

	returnQuery := `
		RETURN u.uid AS Uid,
		u.userId AS UserId,
		u.username AS Username,
		u.email AS Email,
		u.birthDay AS BirthDay,
		u.birthYear AS BirthYear,
		u.birthMonth AS BirthMonth,
		u.profilePic AS ProfilePic,
	`

	query := matchQuery + setQuery + returnQuery

	return query, err.Error{}
}