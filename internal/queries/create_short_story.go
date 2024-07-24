package queries

import (
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/utils"
)

func CreateShortStoryQuery(genres []string) (string, err.Error) {
	creatorLabel, cErr := utils.GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	writingLabel, wErr := utils.GetNodeLabel("Writing")
	if wErr.E != nil {
		return "", wErr
	}

	genreLabels, gErr := BuildGenreLabels(genres)
	if gErr.E != nil {
		return "", gErr
	}

	createdLabel, lErr := utils.GetRelationshipLabel("CREATED")
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
		w.thumbnail AS thumbnail,
		w.uid AS neoId,
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
