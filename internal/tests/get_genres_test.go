package test

import (
	"testing"
	"freecreate/internal/utils"
)

func TestGetGenres(t *testing.T) {
	genres := []string{"Action", "Adventure", "Comedy", "Drama", "HistoricalFiction", "Horror", "Fantasy", "LiteraryFiction", "MagicalRealism", "Mystery", "Realism", "Romance", "SliceOfLife", "SocialFiction", "Superhero", "Thriller"}
	genreMap := make(map[string]bool)

	for _, genre := range genres {
		genreMap[genre] = false
	}

	generatedGenres := utils.GetGenres()

	for _, generatedGenre := range generatedGenres {
		match := false
		for _, genre := range genres {
			if generatedGenre == genre{
				genreMap[genre] = true
				match = true
				break
			}
		}

		if !match {
			t.Errorf("%s from result does not exist in test case", generatedGenre)
		}
	}

	for key, val := range genreMap{
		if !val{
			t.Errorf("%s missing from result", key)
		}
	}
}