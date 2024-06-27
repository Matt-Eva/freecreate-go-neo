package tests

import (
	"freecreate/internal/utils"
	"testing"
)

func TestGetTimeFrames(t *testing.T) {
	timeFrames := []string{"mostRecent", "pastDay", "pastWeek", "pastMonth", "pastYear", "allTime"}
	result := utils.GetTimeFrames()
	frameMap := make(map[string]bool)

	for _, frame := range timeFrames {
		frameMap[frame] = false
	}

	for _, timeFrame := range result {
		err := true
		for _, frame := range timeFrames {
			if timeFrame == frame {
				err = false
				frameMap[frame] = true
			}
		}

		if err {
			t.Errorf("Time frame '%s' from result not present in test case", timeFrame)
		}
	}

	for frame, present := range frameMap {
		if !present {
			t.Errorf("Time frame '%s' from test case not present in result", frame)
		}
	}
}
