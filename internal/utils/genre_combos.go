package utils

import (
	"fmt"
	"slices"
)

func GenerateGenreCombos () [][]string {
	genres := GetGenres()
	genreMap := make(map[string]bool)
	genreCombos := make([][]string, 0, 696)

	for _, genre := range genres{
		genreCombos = append(genreCombos, []string{genre})
	}

	for _, genre := range genres {
		for _, g := range genres {
			sorted := []string{genre, g}
			slices.Sort(sorted)
			combo := fmt.Sprintf("%s:%s", sorted[0], sorted[1])
			if g != genre && !genreMap[combo] {
				genreCombos = append(genreCombos, sorted)
				genreMap[combo] = true
			}
		}
	}

	for _, genre := range genres{
		for _, gen := range genres{
			for _, g := range genres{
				sorted := []string{genre, gen, g}
				slices.Sort(sorted)
				combo := fmt.Sprintf("%s:%s:%s", sorted[0],sorted[1], sorted[2])
				if genre != gen && genre != g && gen != g && !genreMap[combo]{
					genreCombos = append(genreCombos, sorted)
					genreMap[combo] = true
				} 
			}
		}
	}

	return genreCombos
}

// to be used in test file
// needs to be tested itself

// we will test our various genre combos using combination math, rather than
// writing things out by hand (we could have just done that in the first place)
// reference for formula and underlying concepts:
// https://www.mathsisfun.com/combinatorics/combinations-permutations.html
// formula C(n,r) = n!/((n-r)! * r!)
// where n represents the total number of items and r represents the number of
// items within each combo
// Note: 0! = 1 for simplicity sake
// For our purposes, we want to run C(len(genres), 1) + C(len(genres), 2) + C(len(genres), 3)
// This gives us the total number of combinations
// We know that each genre will appear an equal number of times across the combinations
// To calculate the number of times it appears, run C(n-1, r-1)
// Rationale behind formula below.

func CalculateGenreCombos() int{
	genres := GetGenres()

	numberOfGenres := len(genres)
	totalGenreNumFactorial := 1

	for i:= 1; i <= numberOfGenres; i ++{
		totalGenreNumFactorial = totalGenreNumFactorial * i
	}

	comboCount := 0 // should be 696 // that means we will be running 696 * 5 queries to cache frequent search data!
	maxCombo := 3

	for comboNum := 1; comboNum <= maxCombo; comboNum ++ {
		comboFactorial := 1
		for i := 1; i <= comboNum; i++{
			comboFactorial = comboFactorial * i
		}
		
		genreNumLessComboFactorial := 1
		for i:= 1; i<= numberOfGenres - comboNum; i++{
			
			genreNumLessComboFactorial = genreNumLessComboFactorial * i
		}
		comboCount += (totalGenreNumFactorial / (comboFactorial * genreNumLessComboFactorial))
	}

	return comboCount
}

func CalculateGenreAppearances() int{
	genres := GetGenres()
	numberOfGenres := len(genres)
	maxCombo := 3

	numberOfGenresLessOne := numberOfGenres - 1
	totalGenreNumFactorialLessOne := 1

	for i:= 1; i <= numberOfGenresLessOne; i++ {
		totalGenreNumFactorialLessOne = totalGenreNumFactorialLessOne * i
	}

	appearanceComboCount := 1 // to account for case of C(15, 0) with C(16, 1)
	maxComboLessOne := maxCombo - 1

	for comboNum := 1; comboNum <= maxComboLessOne; comboNum ++ {
		comboFactorial := 1
		for i := 1; i <= comboNum; i++{
			comboFactorial = comboFactorial * i
		}
		
		genreNumLessComboFactorial := 1
		for i:= 1; i<= numberOfGenresLessOne - comboNum; i++{
			
			genreNumLessComboFactorial = genreNumLessComboFactorial * i
		}

		appearanceComboCount += (totalGenreNumFactorialLessOne / (comboFactorial * genreNumLessComboFactorial))
	}

	return appearanceComboCount
}