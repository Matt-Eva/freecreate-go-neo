package handlers

import (
	"freecreate/internal/validators"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type RecentQueryStruct struct {
}

func SearchMostRecentHandler(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext) {
	params := r.URL.Query()
	validatedParams, err := validators.ValidateSearchParams(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	BuildMostRecentQuery(validatedParams)
}

func BuildMostRecentQuery(paramStruct validators.ParamStruct) (RecentQueryStruct, error) {
	return RecentQueryStruct{}, nil
}
