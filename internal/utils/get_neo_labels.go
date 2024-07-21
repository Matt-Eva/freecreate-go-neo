package utils

func GetGenres() []string {
	return []string{"Action", "Adventure", "Comedy", "Drama", "HistoricalFiction", "Horror", "Fantasy", "LiteraryFiction", "MagicalRealism", "Mystery", "Realism", "Romance", "ScienceFiction", "SliceOfLife", "SocialFiction", "SpeculativeFiction", "Superhero", "Thriller"}
}

func GetNodeLabels() map[string]string {
	return map[string]string{
		"Writing": "Writing",
		"User":    "User",
		"Creator": "Creator",
	}
}

func GetRelationshipLabels() map[string]string {
	return map[string]string{
		"IS_CREATOR": "IS_CREATOR",
	}
}
