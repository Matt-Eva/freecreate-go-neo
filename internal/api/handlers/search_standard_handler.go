package handlers

import (
	"errors"
	"fmt"
	"freecreate/internal/utils"
	"net/http"
	"net/url"
)

type ParamStruct struct {
	TimeFrame string
	WritingType string
	Genres []string
	Tags []string
	Name string
}

type QueryStruct struct {
	RankQuery string
	RelRankQuery string
	QueryParams string
}

type Results struct {
	RankedResults []string `json:"rankedResults"`
	RelRankedResults []string `json:"relRankedResults"`
}

func SearchStandardHandler(w http.ResponseWriter, r *http.Request, neo string) {
	params := r.URL.Query()

	paramStruct, paramErr := ValidateSearchParams(params)
	if paramErr != nil{
		http.Error(w, paramErr.Error(), http.StatusUnprocessableEntity)
		return
	}

	queryStruct, buildErr := BuildSearchQuery(paramStruct)
	if buildErr != nil {
		http.Error(w, buildErr.Error(), http.StatusUnprocessableEntity)
		return
	}

	RunQuery(queryStruct)

}

func ValidateSearchParams( params url.Values) (ParamStruct, error) {
	timeFrames := utils.GetTimeFrames()
	var paramStruct ParamStruct

	if len(params["timeFrame"]) == 0{
		return ParamStruct{}, errors.New("no time frame specified")
	} else if !timeFrames[params["timeFrame"][0]]{
		errorMsg := fmt.Sprintf("Time frame '%s' is not a valid time frame", params["timeFrame"][0])
		return ParamStruct{}, errors.New(errorMsg)
	} else {
		paramStruct.TimeFrame = params["timeFrame"][0]
	}

	if len(params["writingType"]) == 0 {
		paramStruct.WritingType = ""
	} else {
		writingType, err := utils.ValidateWritingType(params["writingType"][0])
		if err != nil {
			return paramStruct, err
		}
		paramStruct.WritingType = writingType
	}

	genres := params["genres"]
	validatedGenres, err := utils.ValidateGenres(genres)
	if err != nil{
		return ParamStruct{}, err
	}
	paramStruct.Genres = validatedGenres

	if len(params["name"]) == 0{
		paramStruct.Name = ""
	} else {
		paramStruct.Name = params["name"][0]
	}

	paramStruct.Tags = params["tags"]

	return paramStruct, nil
}

func BuildSearchQuery(ParamStruct ParamStruct) (QueryStruct, error) {
	var queryStruct QueryStruct
	var err error

	if ParamStruct.TimeFrame == "mostRecent"{
		queryStruct, err = BuildMostRecentQuery(ParamStruct)
	} else if ParamStruct.TimeFrame == "allTime"{
		queryStruct, err = BuildAllTimeQuery(ParamStruct)
	} else {
		queryStruct, err = BuildStandardQuery(ParamStruct)
	}

	return queryStruct, err
}

func BuildMostRecentQuery(ParamStruct ParamStruct) (QueryStruct, error) {

}

func BuildAllTimeQuery(ParamStruct ParamStruct) (QueryStruct, error) {

}

func BuildStandardQuery(ParamStruct ParamStruct) (QueryStruct, error) {
	var queryStruct QueryStruct

	timeFrame, err := utils.CalculateTimeFrame(ParamStruct.TimeFrame)
	if err != nil {
		return queryStruct, err
	}

	// this doesn't account for novel and collection updates
	timeFrameQuery := fmt.Sprintf("WHERE %d < w.created_at < %d")
	
	queryLabels := "w:Writing"
	

	for _, genre := range ParamStruct.Genres{
		genreLabel := fmt.Sprintf(":%s", genre)
		queryLabels += genreLabel
	}

	
}

func RunQuery(queryStruct QueryStruct){
	if (queryStruct.RankQuery == queryStruct.RelRankQuery){

	} else {

	}
}