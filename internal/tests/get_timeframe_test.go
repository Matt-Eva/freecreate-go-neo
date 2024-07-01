package tests

import (
	"freecreate/internal/utils"
	"testing"
)

func TestGetTimeFrames(t *testing.T) {
	timeFrames := map[string]bool{"mostRecent": true, "pastDay": true, "pastWeek": true, "pastMonth": true, "pastYear": true, "allTime": true}
	isPresent := make(map[string]bool)

	for key := range timeFrames {
		isPresent[key] = false
	}

	result := utils.GetTimeFrames()


	for key := range result {
		_, ok := timeFrames[key]
		if !ok {
			t.Errorf("time frame '%s' is not present in test case", key)
		}
		isPresent[key] = ok
	}

	for key, present := range isPresent {
		if !present {
			t.Errorf("time frame '%s' should be present in result but is not", key)
		}
	}
}
