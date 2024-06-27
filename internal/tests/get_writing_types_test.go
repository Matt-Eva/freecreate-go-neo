package tests

import (
	"freecreate/internal/utils"
	"testing"
)

func TestGetWritingTypes(t *testing.T) {
	caseWritingTypes := map[string]bool{"shortStory": true, "novelette": true, "novella": true, "novel": true}
	writingTypes := utils.GetWritingTypes()

	for _, writingType := range writingTypes {
		if !caseWritingTypes[writingType] {
			t.Errorf("%s writing type not present in test case", writingType)
		}
		caseWritingTypes[writingType] = false
	}

	for key, val := range caseWritingTypes {
		if val {
			t.Errorf("%s writing type was not present in generated types", key)
		}
	}
}
