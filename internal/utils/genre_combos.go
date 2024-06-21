package utils

import (
	"fmt"
)

func GenerateGenreCombos () [][]string {
	genres := GetGenres()
	genreMap := make(map[string]bool)
	genreCombos := make([][]string)

	for _, genre := range genres{
		genreCombos = append(genreCombos, []string{genre})
	}

	for _, genre := genres {
		for _, g := genres {
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
					genreCombos := append(genreCombos, sorted)
					genreMap[combo] = true
				} 
			}
		}
	}

	return genreCombos
}