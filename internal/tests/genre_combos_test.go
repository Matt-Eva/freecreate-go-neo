package tests

import (
	"freecreate/internal/utils"
	"testing"
)

// we will test our various genre combos using combination math, rather than
// writing things out by hand (we could have just done that in the first place)
// reference for formula and underlying concepts:
// https://www.mathsisfun.com/combinatorics/combinations-permutations.html
// formula C(n,r) = n!/((n-r)! * r!)
// where n represents the total number of items and r represents the number of
// items within each combo
// Note: 0! = 1 for simplicity sake
// For our purposes, we want to run C(len(genres), 1) + C(len(genres), 2) + C(len(genres), 3)
// then divide it by len(genres) to get the number of times each genre appears
// Rationale behind formula below.

func TestGenerateGenreCombos(t *testing.T){
	genres := utils.GetGenres()
	genreMap := make(map[string]int)

	for _, genre := range genres {
		genreMap[genre] = 0
	}

	generatedGenres := utils.GenerateGenreCombos()
	comboCount := utils.CalculateGenreCombos()
	
	if len(generatedGenres) != comboCount{
		t.Errorf("Number of generated genres is %d, when it should be %d", len(generatedGenres), comboCount)
	}

	for _, slice := range generatedGenres{
		for _, genre := range slice{
			genreMap[genre] += 1
		}
	}

	genreAppearances := utils.CalculateGenreAppearances()
	for key, val := range genreMap{
		if val != genreAppearances{
			t.Errorf("Count for genre '%s' does not match %d: instead is %d", key, genreAppearances, val)
		}
	}
}

func TestGenreCombos(t *testing.T){
	comboCount := utils.CalculateGenreCombos()
	if comboCount != 696 {
		t.Errorf("combo count does not match 696; instead == %d", comboCount )
	}
}

func TestGenreAppearances(t *testing.T) {
	appearanceCount := utils.CalculateGenreAppearances()
	if appearanceCount != 121{
		t.Errorf("appearance count does not match 120; instead == %d", appearanceCount)
	}
}