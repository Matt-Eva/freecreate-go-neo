package validators

import (
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/queries"
	"freecreate/internal/utils"
	"net/url"
)

type ParamStruct struct {
	TimeFrame   string
	WritingType string
	Genres      []string
	Tags        []string
	Name        string
}

func ValidateSearchParams(params url.Values) (ParamStruct, err.Error) {
	timeFrames := utils.GetTimeFrames()
	var paramStruct ParamStruct

	if len(params["timeFrame"]) == 0 {
		return ParamStruct{}, err.New("no time frame specified")
	} else if !timeFrames[params["timeFrame"][0]] {
		errorMsg := fmt.Sprintf("Time frame '%s' is not a valid time frame", params["timeFrame"][0])
		return ParamStruct{}, err.New(errorMsg)
	} else {
		paramStruct.TimeFrame = params["timeFrame"][0]
	}

	if len(params["writingType"]) == 0 || params["writingType"][0] == "any" {
		paramStruct.WritingType = ""
	} else {
		writingType, wErr := ValidateWritingType(params["writingType"][0])
		if wErr.E != nil {
			return paramStruct, wErr
		}
		paramStruct.WritingType = writingType
	}

	genres := params["genres"]
	validatedGenres, vErr := queries.ValidateGenreLabels(genres)
	if vErr.E != nil {
		return ParamStruct{}, vErr
	}
	paramStruct.Genres = validatedGenres

	if len(params["name"]) == 0 {
		paramStruct.Name = ""
	} else {
		paramStruct.Name = params["name"][0]
	}

	paramStruct.Tags = params["tags"]

	return paramStruct, err.Error{}
}
