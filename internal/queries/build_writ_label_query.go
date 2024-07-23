package queries

import (
	"fmt"
	"freecreate/internal/validators"
	"slices"
)

func BuildWritLabelQuery(genres []string) (string, error) {
	if len(genres) == 0 {
		return "", nil
	}

	validated, err := validators.ValidateGenreLabels(genres)
	if err != nil {
		return "", err
	}

	slices.Sort(validated)

	// need to sort for exact match on mongo cache query
	genreLabels := ""

	for _, genre := range validated {
		genreLabel := fmt.Sprintf(":%s", genre)
		genreLabels += genreLabel
	}

	labels := "(w:Writing" + genreLabels + ")"
	return labels, nil
}
