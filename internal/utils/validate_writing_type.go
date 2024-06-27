package utils

import (
	"errors"
	"fmt"
)

func ValidateWritingType(writingType string) (string, error) {
	types := GetWritingTypes()
	for _, t := range types {
		if writingType == t {
			return t, nil
		}
	}

	errorMessage := fmt.Sprintf("%s is not a valid writing type", writingType)
	return "", errors.New(errorMessage)
}
