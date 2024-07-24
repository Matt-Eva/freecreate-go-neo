package queries

import (
	"strings"
	"testing"
)

func TestCreateShortStoryQuery(t *testing.T) {
	validQuery := `
		MATCH (c:Creator {uid: $creatorId})
		CREATE (w:Writing:Fantasy $shortStoryParams) <- [r:CREATED] - (c)
		RETURN c.name AS author,
		c.profilePic AS authorImg,
		c.creatorId AS authorId,
		w AS shortStory,
		type(r) AS relationship
	`
	generatedQuery, e := CreateShortStoryQuery([]string{"Fantasy"})
	if e.E != nil {
		e.Log()
		t.Fatalf("error")
	}

	strippedValid := strings.ReplaceAll(validQuery, " ", "")
	strippedGenerated := strings.ReplaceAll(generatedQuery, " ", "")

	if strippedValid != strippedGenerated {
		t.Fatalf("generated query '%s' does not match valid query '%s'", strippedGenerated, strippedValid)
	}
}

// func TestCreateShortStoryParams(t *testing.T){
// 	validParams := map[string]any{
// 		"creatorId": "1",
// 		"writingParams": map[string]any{
// 			"title": "test",
// 			"description": "test",
// 			"creatorId": "1",
// 		},
// 	}
// 	p := models.PostedWriting{
// 		Title: "test",
// 		Description: "test",
// 		WritingType: "shortStory",
// 		Thumbnail: "",
// 		CreatorId: "1",
// 	}
// 	year := 2024
// 	shortStory, vErr := models.MakeShortStory(p, year)
// 	if vErr.E != nil {
// 		vErr.Log()
// 		t.Errorf("above error occured")
// 	}
// }
