package tests

import (
	"freecreate/internal/utils"
	"testing"
)

func TestGetGenres(t *testing.T) {
	genreMap := map[string]bool{
		"Action": false,
		"Adventure": false, 
		"Comedy": false, 
		"Drama": false, 
		"HistoricalFiction": false, 
		"Horror": false, 
		"Fantasy": false, 
		"LiteraryFiction": false, 
		"MagicalRealism": false, 
		"Mystery": false, 
		"Realism": false, 
		"Romance": false, 
		"SliceOfLife": false, 
		"SocialFiction": false, 
		"Superhero": false, 
		"Thriller": false,
	}

	genres := utils.GetGenres()

	for _, genre := range genres {
		_, ok := genreMap[genre]
		if ok {
			genreMap[genre] = true
		} else {
			t.Errorf("'%s' genre from generated genres not present in test case", genre)
		}
	}

	for key, val := range genreMap {
		if !val {
			t.Errorf("'%s' genre missing from generated genres", key)
		}
	}
}
