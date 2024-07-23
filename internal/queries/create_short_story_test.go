package queries

import (
	"strings"
	"testing"
)

func TestCreateShortStoryQuery(t *testing.T){
	validQuery := `
		MATCH (c:Creator {uid: $creatorId})
		CREATE (w:Writing $shortStoryParams) <- [r:CREATED] - (c)
		RETURN c.name AS author,
		c.profilePic AS authorImg,
		w.title AS title,
		w.description AS description,
		w.thumbnail AS thumbnail,
		type(r) AS relationship
	`
	generatedQuery, e := CreateShortStoryQuery()
	if e.E != nil {
		e.Log()
		t.Fatalf("error")
	}
	
	strippedValid := strings.ReplaceAll(validQuery, " ", "")
	strippedGenerated := strings.ReplaceAll(generatedQuery, " ", "")
	
	if strippedValid != strippedGenerated{
		t.Fatalf("generated query '%s' does not match valid query '%s'", strippedGenerated, strippedValid)
	}
}