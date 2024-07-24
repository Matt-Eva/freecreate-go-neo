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
	`
	generatedQuery, e := CreateShortStoryQuery([]string{"Fantasy"}, []string{""})
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
