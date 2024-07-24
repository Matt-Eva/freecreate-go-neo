package queries

import (
	"fmt"
	"freecreate/internal/err"
	"slices"
)

func BuildGenreLabels(genres []string) (string, err.Error){
	validatedGenres, vErr := validateGenreLabels(genres)
	if vErr.E != nil {
		return "", vErr
	}

	slices.Sort(validatedGenres)

	// need to sort for exact match on mongo cache query
	genreLabels := ""

	for _, genre := range validatedGenres {
		genreLabel := fmt.Sprintf(":%s", genre)
		genreLabels += genreLabel
	}

	return genreLabels, err.Error{}
}


func validateGenreLabels(genreLabels []string) ([]string, err.Error) {
	genres := GetGenres()
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
