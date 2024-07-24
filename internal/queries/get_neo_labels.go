package queries

import (
	"fmt"
	"freecreate/internal/err"
)

func GetGenres() []string {
	return []string{"Action", "Adventure", "Comedy", "Drama", "HistoricalFiction", "Horror", "Fantasy", "LiteraryFiction", "MagicalRealism", "Mystery", "Realism", "Romance", "ScienceFiction", "SliceOfLife", "SocialFiction", "SpeculativeFiction", "Superhero", "Thriller"}
}

var nodeLabelMap = map[string]string{
	"Writing": "Writing",
	"User":    "User",
	"Creator": "Creator",
}

func GetNodeLabel(label string) (string, err.Error) {
	l, ok := nodeLabelMap[label]
	if !ok {
		msg := fmt.Sprintf("label '%s' is not a valid label", label)
		return "", err.New(msg)
	}

	return l, err.Error{}
}

var relLabelMap = map[string]string{
	"IS_CREATOR": "IS_CREATOR", // User -> Creator
	"CREATED":    "CREATED",    // Creator -> Writing
}

func GetRelationshipLabel(label string) (string, err.Error) {
	l, ok := relLabelMap[label]
	if !ok {
		msg := fmt.Sprintf("label '%s' is not a valid label", label)
		return "", err.New(msg)
	}

	return l, err.Error{}
}