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

type CreatedUser struct {
	Uid         string
	UniqueName string
	Username    string
	Email       string
	ProfilePic  string
	BirthYear   int
	BirthMonth  int
	BirthDay    int
	CreatedAt   int64
	UpdatedAt   int64
}

func CreateUser(ctx context.Context, neo neo4j.DriverWithContext, user models.User) (CreatedUser, err.Error) {
	_, uErr := checkUniqueUser(ctx, neo, user.UniqueName)
	if uErr.E != nil {
		return CreatedUser{},  uErr
	}
	
	params := utils.StructToMap(user)
	query, qErr := buildCreateUserQuery()
	if qErr.E != nil {
		return CreatedUser{}, qErr
	}

	db := os.Getenv("NEO_DB")
	if db == "" {
		return CreatedUser{}, err.New("could not get neo db env variable")
	}
	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return CreatedUser{}, e
	}
	if len(result.Records) < 1 {
		return CreatedUser{}, err.New("db query returned no records")
	}

	resultMap := result.Records[0].AsMap()
	var createdUser CreatedUser
	if e := utils.MapToStruct(resultMap, &createdUser); e.E != nil {
		return CreatedUser{}, e
	}

	return createdUser, err.Error{}
}

func checkUniqueUser(ctx context.Context, neo neo4j.DriverWithContext, uniqueName string)(int, err.Error){
	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return 500, uErr
	}

	query := fmt.Sprintf("MATCH (u:%s {uniqueName: $uniqueName}) RETURN u.uniqueName AS uniqueName", userLabel)
	params := map[string]any {
		"uniqueName": uniqueName,
	}

	db := os.Getenv("NEO_DB")
	if db == ""{
		return 500, err.New("neodb environment variable is empty")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return 500, err.NewFromErr(nErr)
	}
	if len(result.Records) > 0{
		return 422, err.New("user id already in use")
	}

	return 200, err.Error{}
}

func buildCreateUserQuery() (string, err.Error) {
	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	createQuery := fmt.Sprintf("CREATE(u:%s $userParams)", userLabel)

	returnQuery := `
		RETURN u.uid AS Uid,
		u.uniqueName AS UniqueName,
		u.username AS Username,
		u.email AS Email,
		e.profilePic AS ProfilePic,
		u.birthYear AS BirthYear,
		u.birthMonth AS BirthMonth,
		u.birthDay AS BirthDay,
		u.createdAt AS CreatedAt,
		u.updatedAt AS UpdatedAt
	`

	query := createQuery + returnQuery

	return query, err.Error{}
}
