package seeds

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedShortStories(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client) err.Error {
	creators, cErr := getCreators(ctx, neo)
	if cErr.E != nil {
		return cErr
	}

	for _, creatorId := range creators{
		shortStory, mErr := makeShortStory(creatorId)
		if mErr.E != nil {
			return mErr
		}
		sErr := seedShortStory(ctx, neo, shortStory)
		if sErr.E != nil {
			return sErr
		}
	}
	return err.Error{}
}

func getCreators(ctx context.Context, neo neo4j.DriverWithContext)([]string, err.Error){
	creators := make([]string, 0)

	query := `
		MATCH (c:Creator)
		WHERE c.seed = true
		RETURN c.uid AS uid
	`

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return []string{}, e
	}

	if len(result.Records) == 0 {
		e := err.New("no records returned for get seed creators")
		return []string{}, e
	}

	for _, record := range result.Records{
		uid, ok := record.Get("uid")
		if !ok {
			e := err.New("creator seed record does not have uid attribute")
			return []string{},e
		}

		creatorId, ok := uid.(string)
		if !ok {
			e := err.New("creator seed field uid could not be converted to string")
			return []string{},e
		}
		creators = append(creators, creatorId)
	}

	return creators, err.Error{}
}

func makeShortStory(creatorId string)(models.ShortStory, err.Error){
	year := time.Now().Year()
	p := models.PostedWriting{
		Title: faker.Sentence(),
		Description: faker.Paragraph(),
		WritingType: "shortStory",
		Thumbnail: "",
		CreatorId: creatorId,
	}

	s, mErr := models.MakeShortStory(p, year)
	if mErr.E != nil {
		return models.ShortStory{}, mErr
	}

	return s, err.Error{}
}

func seedShortStory(ctx context.Context, neo neo4j.DriverWithContext, shortStory models.ShortStory) err.Error{
	params := queries.CreateShortStoryParams(shortStory)

	query, qErr := queries.CreateShortStoryQuery()
	if qErr.E != nil {
		return qErr
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return e
	}
	
	if (len(result.Records) < 1){
		e := err.New("no record return from create seed short story")
		return e
	}

	recordMap := result.Records[0].AsMap()

	fmt.Println(recordMap)
	
	return err.Error{}
}
