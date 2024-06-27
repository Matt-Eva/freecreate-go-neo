package utils

import (
	"slices"
	"strings"
)

func BuildFicGenreLabel(genres []string) (string, error) {
	if len(genres) == 0 {
		return "", nil
	}

	validated, err := ValidateGenres(genres)
	if err != nil {
		return "", err
	}

	// need to sort for exact match on redis cache query
	slices.Sort(validated)

	labels := ":" + strings.Join(validated, ":")

	return labels, nil

}
