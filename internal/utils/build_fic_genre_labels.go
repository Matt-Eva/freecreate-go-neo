package utils

import "strings"

func BuildFicGenreLabel(genres []string) (string, error){
	if len(genres) == 0{
		return "", nil
	}

	validated, err := ValidateGenres(genres)
	if err != nil{
		return "", err
	}

	labels := ":" + strings.Join(validated, ":")

	return labels, nil

}