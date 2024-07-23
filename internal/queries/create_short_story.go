package queries

import (
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/utils"
)

func CreateShortStoryQuery() (string, err.Error) {
	creatorLabel, cErr := utils.GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	writingLabel, wErr := utils.GetNodeLabel("Writing")
	if wErr.E != nil {
		return "", wErr
	}

	createdLabel, lErr := utils.GetRelationshipLabel("CREATED")
	if lErr.E != nil {
		return "", lErr
	}

	query := fmt.Sprintf(`
		MATCH (c:%s {uid: $creatorId})
		CREATE (w:%s $shortStoryParams) <-[r:%s] - (c)
		RETURN c.name AS author, 
		c.profilePic AS authorImg,
		c.creatorId AS authorId,  
		w.title AS title, 
		w.description AS description,
		w.thumbnail AS thumbnail,
		type(r) AS relationship
	`, creatorLabel, writingLabel, createdLabel)

	return query, err.Error{}
}

func CreateShortStoryParams(shortStory models.ShortStory) map[string]any {
	params := map[string]any{
		"creatorId":        shortStory.CreatorId,
		"shortStoryParams": utils.NeoParamsFromStruct(shortStory),
	}

	return params
}
