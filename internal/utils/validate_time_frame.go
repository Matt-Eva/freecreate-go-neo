package utils

import (
	"errors"
	"fmt"
)

func ValidateTimeFrame(timeFrame string) (string, error) {
	timeFrames := GetTimeFrames()

	for _, frame := range timeFrames {
		if timeFrame == frame {
			return frame, nil
		}
	}

	errorMessage := fmt.Sprintf("Time frame %s is not a valid time frame", timeFrame)

	return "", errors.New(errorMessage)
}
