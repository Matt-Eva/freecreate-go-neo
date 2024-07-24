package validators

import (
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
)

func ValidateWritingType(writingType string) (string, err.Error) {
	types := utils.GetWritingTypes()
	for _, t := range types {
		if writingType == t {
			return t, err.Error{}
		}
	}

	errorMessage := fmt.Sprintf("%s is not a valid writing type", writingType)
	return "", err.New(errorMessage)
}
