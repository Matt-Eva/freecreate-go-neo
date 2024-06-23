package databases

import (
	"freecreate/internal/utils"
	"strings"
)

func NeoSimpleSearch(timeFrame string, genres [] string){
	simpleSearchQuery := GenerateNeoSimpleSearchQuery(timeFrame, genres)
}

func GenerateNeoSimpleSearchQuery (timeFrame string, genres []string) string {
	validatedGenres := utils.ValidateGenres(genres)
	genreLabels := strings.Join(validatedGenres, ":")
	
}