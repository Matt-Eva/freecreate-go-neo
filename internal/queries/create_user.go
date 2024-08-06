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
	DisplayName string
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
	if e := utils.MapToStruct(resultMap, createdUser); e.E != nil {
		return CreatedUser{}, e
	}

	return createdUser, err.Error{}
}

func buildCreateUserQuery() (string, err.Error) {
	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	createQuery := fmt.Sprintf("CREATE(u:%s $userParams)", userLabel)

	returnQuery := `
		RETURN u.uid AS Uid,
		u.displayName AS DisplayName,
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
