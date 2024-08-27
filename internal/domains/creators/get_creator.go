package creators

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type RetrievedCreator struct {
	Name       string
	CreatorId  string
	ProfilePic string
	About      string
	Uid        string
}

func GetCreatorQuery(ctx context.Context, neo neo4j.DriverWithContext, creatorId string) (RetrievedCreator, err.Error) {
	params := map[string]any{
		"creatorId": creatorId,
	}

	query, qErr := buildGetCreatorQuery()
	if qErr.E != nil {
		return RetrievedCreator{}, qErr
	}

	db := os.Getenv("NEO_DB")
	if db == "" {
		return RetrievedCreator{}, err.New("could not get db environment variable")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return RetrievedCreator{}, e
	}
	if len(result.Records) < 1 {
		return RetrievedCreator{}, err.New("no records returned from db")
	}

	resultMap := result.Records[0].AsMap()

	var creator RetrievedCreator
	if e := utils.MapToStruct(resultMap, &creator); e.E != nil {
		return RetrievedCreator{}, e
	}

	return creator, err.Error{}
}

func buildGetCreatorQuery() (string, err.Error) {
	creatorLabel, lErr := utils.GetNodeLabel("Creator")
	if lErr.E != nil {
		return "", err.Error{}
	}

	matchQuery := fmt.Sprintf("MATCH (c:%s {uid: $creatorId})", creatorLabel)

	returnQuery := `
		RETURN c.name AS Name,
		c.uid AS Uid,
		c.creatorId AS CreatorId,
		c.about AS About,
		c.profilePic AS ProfilePic
	`

	query := matchQuery + returnQuery

	return query, err.Error{}

}
