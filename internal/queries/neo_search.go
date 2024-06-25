package databases

import (
	"fmt"
	"freecreate/internal/utils"
	"strings"
)

func NeoSearch(timeFrame, writingType string, genres []string) {
	if timeFrame == "Most Recent" {

	} else if timeFrame == "All Time" {

	} else {
		simpleSearchQuery := GenerateNeoSimpleSearchQuery(timeFrame, writingType, genres)
	}
}

func GenerateNeoSimpleSearchQuery(timeFrame, writingType string, genres []string) string {
	validatedGenres := utils.ValidateGenres(genres)
	validatedWritingType := utils.ValidateWritingType(writingType)
	genreLabels := strings.Join(validatedGenres, ":")
	timeFrameStruct := utils.CalculateTimeFrame(timeFrame)

	query := fmt.Sprintf("MATCH (w:Writing%s) WHERE writingType = %s")
	timeQuery := ""

}
