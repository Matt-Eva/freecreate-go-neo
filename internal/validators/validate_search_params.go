package validators

import (
	"errors"
	"fmt"
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

func ValidateSearchParams(params url.Values) (ParamStruct, error) {
	timeFrames := utils.GetTimeFrames()
	var paramStruct ParamStruct

	if len(params["timeFrame"]) == 0 {
		return ParamStruct{}, errors.New("no time frame specified")
	} else if !timeFrames[params["timeFrame"][0]] {
		errorMsg := fmt.Sprintf("Time frame '%s' is not a valid time frame", params["timeFrame"][0])
		return ParamStruct{}, errors.New(errorMsg)
	} else {
		paramStruct.TimeFrame = params["timeFrame"][0]
	}

	if len(params["writingType"]) == 0 || params["writingType"][0] == "any" {
		paramStruct.WritingType = ""
	} else {
		writingType, err := ValidateWritingType(params["writingType"][0])
		if err != nil {
			return paramStruct, err
		}
		paramStruct.WritingType = writingType
	}

	genres := params["genres"]
	validatedGenres, err := ValidateGenreLabels(genres)
	if err != nil {
		return ParamStruct{}, err
	}
	paramStruct.Genres = validatedGenres

	if len(params["name"]) == 0 {
		paramStruct.Name = ""
	} else {
		paramStruct.Name = params["name"][0]
	}

	paramStruct.Tags = params["tags"]

	return paramStruct, nil
}
