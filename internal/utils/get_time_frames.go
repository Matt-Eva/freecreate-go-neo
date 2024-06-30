package utils

func GetTimeFrames() map[string]bool {
	return map[string]bool{"mostRecent": true, "pastDay": true, "pastWeek": true, "pastMonth": true, "pastYear": true, "allTime": true}
}
