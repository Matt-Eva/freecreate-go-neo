package search_handler

import (
	"net/http"
	"net/url"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type RecentQueryStruct struct {
}

func SearchMostRecentHandler(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext) {
	params := r.URL.Query()

	BuildMostRecentQuery(params)
}

func BuildMostRecentQuery(params url.Values) (RecentQueryStruct, error) {
	return RecentQueryStruct{}, nil
}
