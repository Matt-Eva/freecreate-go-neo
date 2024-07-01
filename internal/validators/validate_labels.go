package validators

import (
	"errors"
	"fmt"
	"freecreate/internal/utils"
)

func ValidateGenreLabels(genreLabels []string) ([]string, error) {
	genres := utils.GetGenres()
	validatedLabels := make([]string, 0, 3)
	validatedMap := make(map[string]bool)

	for _, label := range genreLabels {
		validatedMap[label] = false
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
			return []string{}, errors.New(errorMsg)
		}
	}

	return validatedLabels, nil
}