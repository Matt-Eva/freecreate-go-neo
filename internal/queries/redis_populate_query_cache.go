package databases

import (
	"freecreate/internal/utils"
	"strings"
)

func PopulateRedisQueryCache() {
	queryComboMap := utils.AssembleCachePopulationCombos()
	neoQueries := make([]string, 0, 696*5*5)

	for writingType, timeMap := range queryComboMap {
		for timeFrame, genreSlices := range timeMap {
			for _, genreSlice := range genreSlices {
				neoQuery := BuildNeoPopulateCacheQuery(writingType, timeFrame, genreSlice)
				neoQueries = append(neoQueries, neoQuery)
			}
		}
	}

}

func BuildNeoPopulateCacheQuery(writingType, timeFrame string, genres []string) string {
	validatedGenres := utils.ValidateGenres(genres)
	validatedWritingType := utils.ValidateWritingType(writingType)
	genreLabels := strings.Join(validatedGenres, ":")
	timeFrameQuery := ""

	if timeFrame == "Most Recent" {

	} else if timeFrame == "All Time" {

	}

	return ""
}

func BuildRedisPopulateCacheQuery() {

}
