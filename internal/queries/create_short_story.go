package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateShortStory(ctx context.Context, neo neo4j.DriverWithContext, shortStory models.ShortStory){}

func CreateShortStoryQuery(genres []string) (string, err.Error) {
	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	writingLabel, wErr := GetNodeLabel("Writing")
	if wErr.E != nil {
		return "", wErr
	}

	genreLabels, gErr := BuildGenreLabels(genres)
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
		w AS shortStory,
		type(r) AS relationship
	`, creatorLabel, writingLabel, genreLabels, createdLabel)

	return query, err.Error{}
}

func CreateShortStoryParams(shortStory models.ShortStory) map[string]any {
	params := map[string]any{
		"creatorId":        shortStory.CreatorId,
		"shortStoryParams": NeoParamsFromStruct(shortStory),
	}

	return params
}
