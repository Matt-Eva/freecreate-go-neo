package validators

import (
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
)

func ValidateGenreLabels(genreLabels []string) ([]string, err.Error) {
	genres := utils.GetGenres()
	validatedLabels := make([]string, 0, 3)
	validatedMap := make(map[string]bool)

	for _, label := range genreLabels {
		_, ok := validatedMap[label]
		if ok {
			continue
		}
		
		for _, genre := range genres {
			if label == genre {
				validatedLabels = append(validatedLabels, genre)
				validatedMap[label] = true
				break
			}
		}
	}

	for key, present := range validatedMap {
		if !present {
			errorMsg := fmt.Sprintf("%s is not a valid genre", key)
			return []string{}, err.New(errorMsg)
		}
	}

	return validatedLabels, err.Error{}
}
