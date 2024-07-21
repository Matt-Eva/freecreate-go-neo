package seeds

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"

	"github.com/go-faker/faker/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedCreators(ctx context.Context, neo neo4j.DriverWithContext) err.Error {
	fmt.Println("seeding creators")
	result, uErr := getSeedUsers(ctx, neo)
	if uErr.E != nil {
		return uErr
	}

	for _, uid := range result {
		sErr := seedCreator(ctx, neo, uid)
		if sErr.E != nil {
			return sErr
		}
	}

	fmt.Println("creators seeded")
	return err.Error{}
}

func getSeedUsers(ctx context.Context, neo neo4j.DriverWithContext) ([]string, err.Error) {
	users := make([]string, 0)

	query := `
		MATCH (u:User)
		WHERE u.seed = true
		RETURN u.uid AS uid
	`

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return []string{}, e
	}
	if len(result.Records) < 1 {
		e := err.New("no seed users in database")
		return []string{}, e
	}

	for _, record := range result.Records {
		uid, ok := record.Get("uid")
		if !ok {
			e := err.New("returned user record does not have uid attribute")
			return []string{}, e
		}

		userId, ok := uid.(string)
		if !ok {
			e := err.New("could not convert user uid to type string")
			return []string{}, e
		}

		users = append(users, userId)
	}

	return users, err.Error{}
}

func seedCreator(ctx context.Context, neo neo4j.DriverWithContext, userId string) err.Error {
	params, gErr := makeSeedCreator(userId)
	if gErr.E != nil {
		return gErr
	}

	query := `
		MATCH (u:User {uid: $userId})	
		CREATE (c:Creator $creatorParams)
		SET c.seed = true
		CREATE (u) - [r:IS_CREATOR] -> (c)
		RETURN type(r) AS rel, u.username AS username, c.name AS creator 
	`
	result, qErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if qErr != nil {
		e := err.NewFromErr(qErr)
		return e
	}
	if len(result.Records) < 1 {
		return err.New("no record returned from seed creator")
	}
	// fmt.Println(result.Records[0].AsMap())

	return err.Error{}
}

func makeSeedCreator(userId string) (map[string]any, err.Error) {
	p := models.PostedCreator{
		Name:       faker.Name(),
		ProfilePic: "",
		CreatorId:  faker.Username(),
		About:      faker.Paragraph(),
	}

	c, gErr := p.GenerateCreator(userId)
	if gErr.E != nil {
		return map[string]any{}, gErr
	}

	params := c.NewCreatorParams()

	return params, err.Error{}
}

func DeleteCreatorSeeds(neo neo4j.DriverWithContext, mongo *mongo.Client) err.Error{

	return err.Error{}
}
