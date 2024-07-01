package tests

import (
	"freecreate/internal/api/handlers"
	"freecreate/internal/utils"
	"strings"
	"testing"
)


func TestSearchStandardHandler(t *testing.T){

}

func TestBuildStandardSearchQuery(t *testing.T){
	type TestCase struct {
		Case utils.ParamStruct
		RankQuery string
		RelRankQuery string
		QueryParams map[string]any
	}

	testCases := []TestCase{
		{
			utils.ParamStruct{
				WritingType: "shortStory",
				Name: "",
				TimeFrame: "pastDay",
				Genres: []string{"ScienceFiction"},
				Tags: []string{},
			},
			"MATCH(w:Writing:ScienceFiction) WHERE $start < w.createdAt < $endWITH w MATCH (w) <- [:CREATED] - (c:Creator) <- [:IS_CREATOR] - (u:User) RETURN w.title AS title, w.description AS description, c.name AS author, u.username AS username ORDER BY w.rank LIMIT 100",
			"MATCH(w:Writing:ScienceFiction) WHERE $start < w.createdAt < $endWITH w MATCH (w) <- [:CREATED] - (c:Creator) <- [:IS_CREATOR] - (u:User) RETURN w.title AS title, w.description AS description, c.name AS author, u.username AS username ORDER BY w.relRank LIMIT 100",
			map[string]any{
				"start": 0,
				"end": 0,
			},
		},
	}

	for _, testCase := range testCases{
		result, err := handlers.BuildStandardSearchQuery(testCase.Case)
		if err != nil {
			t.Errorf(err.Error())
		}
		resultRankQuery := strings.ReplaceAll(result.RankQuery, " ", "")
		resultRelRankQuery := strings.ReplaceAll(result.RelRankQuery, " ", "")
		testRankQuery := strings.ReplaceAll(testCase.RankQuery, " ", "")
		testRelRankQuery := strings.ReplaceAll(testCase.RelRankQuery, " ", "")

		if resultRankQuery != testRankQuery{
			t.Errorf("Result query \n '%s' \n not equal to \n '%s'", resultRankQuery, testRankQuery)
		} else if resultRelRankQuery != testRelRankQuery{
			t.Errorf("Result query \n '%s' \n not equal to \n test case query'%s'", resultRelRankQuery, testRelRankQuery)
		}
	}

}