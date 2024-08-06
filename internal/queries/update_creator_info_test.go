package queries

import (
	"strings"
	"testing"
)

func TestBuildUpdateCreatorInfoQuery(t *testing.T){
	params := map[string]any{
		"name": "bob",
	}
	validQuery := "MATCH (c:Creator {uid: $uid}) SET c.name = $name RETURN c.uid AS Uid, c.name AS Name, c.creatorId AS CreatorId, c.about AS About"
	generatedQuery, qErr := buildUpdateCreatorInfoQuery(params)
	if qErr.E != nil {
		qErr.Log()
		t.Fatal("above error occurred during testing")
	}
	
	strippedValid := strings.ReplaceAll(validQuery, " ", "" )
	strippedGenerated := strings.ReplaceAll(generatedQuery, " ", "")
	if strippedGenerated != strippedValid {
		t.Fatalf ("query strings do not match: \n %s \n \n %s", strippedValid, strippedGenerated)
	}
}