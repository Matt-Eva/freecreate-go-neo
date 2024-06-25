package utils

func ValidateGenres(genreLabels []string) []string {
	genres := GetGenres()
	validatedLabels := make([]string, 0, 3)

	for _, label := range genreLabels {
		for _, genre := range genres {
			if label == genre {
				validatedLabels = append(validatedLabels, genre)
				break
			}
		}
	}

	return validatedLabels
}
