package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/utils"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type CreatedShortStory struct {
	Author       string
	AuthorImg    string
	Title        string
	Description  string
	Uid          string
	Thumbnail    string
	WritingType  string
	CreatedAt    string
	UpdatedAt    string
	LibraryCount int64
	Likes        int64
	Views        int64
	Donations    int64
	Rank         int64
	RelRank      int64
	OriginalYear int
}

func CreateShortStory(ctx context.Context, neo neo4j.DriverWithContext, shortStory models.ShortStory, genres, tags []string) (CreatedShortStory, err.Error) {
	query, qErr := CreateShortStoryQuery(genres, tags)
	if qErr.E != nil {
		return CreatedShortStory{}, err.Error{}
	}

	params := CreateShortStoryParams(shortStory)

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if nErr != nil {
		e := err.NewFromErr(nErr)
		return CreatedShortStory{}, e
	}
	if len(result.Records) < 1 {
		return CreatedShortStory{}, err.New("create short story op returning no record")
	} else if len(result.Records) > 1 {
		return CreatedShortStory{}, err.New("create short story op returning more than one record")
	}

	createdStory := CreatedShortStory{}
	// record := result.Records[0]

	return createdStory, err.Error{}
}

func CreateShortStoryQuery(genres []string, tags []string) (string, err.Error) {
	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	writingLabel, wErr := GetNodeLabel("Writing")
	if wErr.E != nil {
		return "", wErr
	}

	genreLabels, gErr := buildGenreLabels(genres)
	if gErr.E != nil {
		return "", gErr
	}

	createdLabel, lErr := GetRelationshipLabel("CREATED")
	if lErr.E != nil {
		return "", lErr
	}

	query := fmt.Sprintf(`
		MATCH (c:%s {uid: $creatorId})
		CREATE (w:%s%s $shortStoryParams) <-[r:%s] - (c)
		RETURN c.name AS author, 
		c.profilePic AS authorImg,
		c.creatorId AS authorId,  
		w.title AS title,
		w.description AS description,
		w.uid AS uid,
		w.thumbnail AS thumbnail,
		w.writingType AS writingType,
		w.createdAt AS createdAt,
		w.updatedAt AS updatedAt,
		w.libraryCount AS libraryCount,
		w.likes AS likes,
		w.views AS views,
		w.donations AS donations,
		w.rank AS rank,
		w.relRank AS relRank,
		w.originalYear AS originalYear,
		type(r) AS relationship
	`, creatorLabel, writingLabel, genreLabels, createdLabel)

	return query, err.Error{}
}

func CreateShortStoryParams(shortStory models.ShortStory) map[string]any {
	params := map[string]any{
		"creatorId":        shortStory.CreatorId,
		"shortStoryParams": utils.StructToMap(shortStory),
	}

	return params
}
