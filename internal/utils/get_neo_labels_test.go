package utils

import "testing"

func TestGetRelationshipLabel(t *testing.T) {
	validRels := []string{"IS_CREATOR"}
	relMap := make(map[string]bool)

	for _, rel := range validRels {
		relMap[rel] = false
	}

	for _, rel := range validRels {
		match, ok := relLabelMap[rel]
		if !ok {
			t.Errorf("generated relationship labels does not contain the '%s' relationship", rel)
		}
		if match != rel {
			t.Errorf("generated relationship labels contain the '%s' relationship, but the value '%s' is wrong", rel, match)
		}
	}

}
