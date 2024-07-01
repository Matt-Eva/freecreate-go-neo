package validators

import (
	"errors"
	"fmt"
	"freecreate/internal/utils"
)

func ValidateTimeFrame(timeFrame string) (string, error) {
	timeFrames := utils.GetTimeFrames()

	if !timeFrames[timeFrame] {
		errorMessage := fmt.Sprintf("Time frame %s is not a valid time frame", timeFrame)

		return "", errors.New(errorMessage)
	}

	return timeFrame, nil

}
