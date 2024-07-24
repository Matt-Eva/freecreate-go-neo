package queries

import (
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/validators"
	"slices"
)

func BuildGenreLabels(genres []string) (string, err.Error){
	validatedGenres, vErr := validators.ValidateGenreLabels(genres)
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