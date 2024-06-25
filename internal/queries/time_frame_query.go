package databases

func TimeFrameQuery (timeFrame string) string {
	timeFrameQuery := ""

	if timeFrame == "Most Recent" {
		timeFrameQuery = "ORDER BY w.created_at"
	} else if timeFrame == "All Time" {
		timeFrameQuery = ""
	}
}