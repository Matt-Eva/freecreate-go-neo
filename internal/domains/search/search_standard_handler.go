package search_handler

import (
	"net/http"
	"net/url"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type QueryStruct struct {
	RankQuery    string
	RelRankQuery string
	QueryParams  map[string]any
}

type Results struct {
	RankedResults    []string `json:"rankedResults"`
	RelRankedResults []string `json:"relRankedResults"`
}

func SearchStandardHandler(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext) {
	params := r.URL.Query()

	queryStruct, buildErr := BuildStandardSearchQuery(params)
	if buildErr != nil {
		http.Error(w, buildErr.Error(), http.StatusUnprocessableEntity)
		return
	}

	RunQuery(queryStruct)

}

func BuildStandardSearchQuery(params url.Values) (QueryStruct, error) {
	var queryStruct QueryStruct
	// queryParams := make(map[string]any)

	// queryLabels, err := queries.BuildWritLabelQuery(paramStruct.Genres)
	// if err != nil {
	// 	return queryStruct, err
	// }

	// timeFrame, err := utils.CalculateTimeFrame(paramStruct.TimeFrame)
	// if err != nil {
	// 	return queryStruct, err
	// }

	// timeFrameQuery := ""
	// if paramStruct.WritingType == "shortStory" || paramStruct.WritingType == "" {
	// 	timeFrameQuery = "WHERE $start < w.createdAt < $end"
	// } else {
	// 	timeFrameQuery = "WHERE $start < w.latestPublication < $end"
	// }
	// queryParams["start"] = timeFrame.Start
	// queryParams["end"] = timeFrame.End

	// nameQuery := ""
	// if paramStruct.Name != "" {
	// 	nameQuery = " AND WHERE w.title = $title"
	// 	queryParams["title"] = paramStruct.Name
	// }

	// tagQuery := ""
	// for i, tag := range paramStruct.Tags {
	// 	paramKey := fmt.Sprintf("tag%d", i)
	// 	queryParams[paramKey] = tag
	// 	query := fmt.Sprintf(" AND WHERE (w) - [:HAS_TAG] -> (t:Tag {tag: $%s})", paramKey)
	// 	tagQuery += query
	// }

	// getAuthor := queries.BuildGetAuthorQuery()

	// returnStatement := queries.BuildNeoWritReturnQuery()

	// rankedOrder := "ORDER BY w.rank"
	// relRankedOrder := "ORDER BY w.relRank"
	// limitQuery := "LIMIT 100"

	// rankedQuery := fmt.Sprintf("MATCH" + queryLabels + timeFrameQuery + nameQuery + tagQuery + getAuthor + returnStatement + rankedOrder + limitQuery)
	// relRankedQuery := fmt.Sprintf("MATCH" + queryLabels + timeFrameQuery + nameQuery + tagQuery + getAuthor + returnStatement + relRankedOrder + limitQuery)

	// queryStruct.QueryParams = queryParams
	// queryStruct.RankQuery = rankedQuery
	// queryStruct.RelRankQuery = relRankedQuery

	return queryStruct, nil
}

func RunQuery(queryStruct QueryStruct) {
	// if (queryStruct.RankQuery == queryStruct.RelRankQuery){

	// } else {

	// }
}
