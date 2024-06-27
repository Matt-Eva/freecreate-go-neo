package utils

import (
	"errors"
	"fmt"
	"time"
)

type TimeFrame struct {
	Start int64
	End   int64
}

func CalculateTimeFrame(timeFrame string) (TimeFrame, error) {
	now := time.Now().UTC().UnixMilli()
	year := now - 31556952000 // this is technically not necessary, since databases will be sharded by year.
	month := now - 2628000000
	week := now - 604800000
	day := now - 86400000

	dateMap := map[string]TimeFrame{
		"pastYear":  {month, year},
		"pastMonth": {week, month},
		"pastWeek":  {day, week},
		"pastDay":   {now, day},
	}

	dateQueryStruct, ok := dateMap[timeFrame]
	if !ok {
		errorMsg := fmt.Sprintf("%s time frame cannot be used in calculate time frame - calculate_time_frame.go", timeFrame)
		return TimeFrame{}, errors.New(errorMsg)
	}

	return dateQueryStruct, nil
}
