package utils

import "time"

type TimeFrame struct {
	Start int64
	End   int64
}

func CalculateTimeFrame(timeFrame string) TimeFrame {
	now := time.Now().UTC().UnixMilli()
	year := now - 31556952000 // this is technically not necessary, since databases will be sharded by year.
	month := now - 2628000000
	week := now - 604800000
	day := now - 86400000

	dateMap := map[string]TimeFrame{
		"Past Year":  {month, year},
		"Past Month": {week, month},
		"Past Week":  {day, week},
		"Past Day":   {now, day},
	}

	dateQueryStruct := dateMap[timeFrame]

	return dateQueryStruct
}
