package seeds

import (
	"context"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreatorData struct {
	CreatorUid string
	CreatorId  string
	UserUid    string
}

func SeedShortStories(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client) err.Error {
	creators, cErr := getCreators(ctx, neo)
	if cErr.E != nil {
		return cErr
	}

	for _, creatorData := range creators {
		shortStory, mErr := makeShortStory(creatorData.CreatorId)
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

func getCreators(ctx context.Context, neo neo4j.DriverWithContext) ([]CreatorData, err.Error) {

	creators := make([]CreatorData, 0)

	query := `
		MATCH (c:Creator) <- [:IS_CREATOR] - (u:User)
		WHERE c.seed = true
		RETURN c.uid AS creatorUid, c.creatorId AS creatorId, u.uid AS userUid
	`

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return creators, e
	}

	if len(result.Records) == 0 {
		e := err.New("no records returned for get seed creators")
		return creators, e
	}

	for _, record := range result.Records {
		creatorData := CreatorData{}

		cUid, ok := record.Get("creatorUid")
		if !ok {
			e := err.New("creator seed record does not have creatorUid attribute")
			return creators, e
		}
		creatorUid, ok := cUid.(string)
		if !ok {
			e := err.New("creator seed field creatorUid could not be converted to string")
			return creators, e
		}
		creatorData.CreatorUid = creatorUid

		cId, ok := record.Get("creatorId")
		if !ok {
			e := err.New("creator seed record does not have creatorId attribute")
			return creators, e
		}
		creatorId, ok := cId.(string)
		if !ok {
			e := err.New("creator seed field creatorId could not be converted to string")
			return creators, e
		}
		creatorData.CreatorId = creatorId

		uUid, ok := record.Get("userUid")
		if !ok {
			e := err.New("creator seed record does not have userUid attribute")
			return creators, e
		}
		userUid, ok := uUid.(string)
		if !ok {
			e := err.New("creator seed field userUid could not be converted to string")
			return creators, e
		}
		creatorData.UserUid = userUid

		creators = append(creators, creatorData)
	}

	return creators, err.Error{}
}

func makeShortStory(creatorId string) (models.ShortStory, err.Error) {
	year := time.Now().Year()
	p := models.PostedWriting{
		Title:       faker.Sentence(),
		Description: faker.Paragraph(),
		WritingType: "shortStory",
		Thumbnail:   "",
		CreatorId:   creatorId,
	}

	s, mErr := models.MakeShortStory(p, year)
	if mErr.E != nil {
		return models.ShortStory{}, mErr
	}

	s.Published = true

	return s, err.Error{}
}

func seedShortStory(ctx context.Context, neo neo4j.DriverWithContext, shortStory models.ShortStory) err.Error {
	params := queries.CreateShortStoryParams(shortStory)

	genres := queries.GetGenres()
	selectedGenres := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		genre := genres[rand.Intn(len(genres))]
		selectedGenres = append(selectedGenres, genre)
	}
	query, qErr := queries.CreateShortStoryQuery(selectedGenres, []string{""})
	if qErr.E != nil {
		return qErr
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return e
	}
	if len(result.Records) < 1 {
		e := err.New("no record return from create seed short story")
		return e
	}

	// recordMap := result.Records[0].AsMap()
	// fmt.Println(recordMap)

	return err.Error{}
}
