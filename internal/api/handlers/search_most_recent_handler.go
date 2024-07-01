package handlers

import (
	"freecreate/internal/utils"
	"net/http"
)

type RecentQueryStruct struct {

}

func SearchMostRecentHandler(w http.ResponseWriter, r *http.Request, neo string) {
	params := r.URL.Query()
	validatedParams, err := utils.ValidateSearchParams(params)
	if err != nil {
		
	}

	BuildMostRecentQuery(validatedParams)
}

func BuildMostRecentQuery(paramStruct utils.ParamStruct) (RecentQueryStruct, error) {
	return RecentQueryStruct{}, nil
}
