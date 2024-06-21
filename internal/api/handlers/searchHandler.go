package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"freecreate/internal/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request, neo, mongo, redis string) {
	params := r.URL.Query()
	detectSearchType(params)
}

func detectSearchType(params url.Values) ([]string, error) {
	if params["searchType"][0] == "writing" {
		buildWritingQuery(params)
	} else if params["searchType"][0] == "writer" {
		buildWriterQuery(params)
	}

	return []string{}, errors.New("")
}

func buildWritingQuery(params url.Values) ([]string, error) {
	date := params["date"][0]
	genres := params["genre"]
	tags := params["tag"]
	title := params["writingTitle"]
	// writingType := params["writingType"]

	genreQuery := buildWritingGenreQuery(genres)
	fmt.Println(genreQuery)

	dateQuery := handleDateQueryType(date)
	fmt.Println(dateQuery)

	if len(tags) == 0 && len(title) == 0 {
		fmt.Println("hit cache")
	} else {
	}

	return []string{}, nil
}

func buildWritingGenreQuery(genres []string) string {
	genreLabels := utils.GetGenres()

	queryLabels := "w:Writing"

	for _, genre := range genres {
		for _, label := range genreLabels {
			if genre == label {
				queryLabels += fmt.Sprintf(":%s", label)
			}
		}
	}

	genreQuery := fmt.Sprintf("MATCH (%s)", queryLabels)
	return genreQuery
}

func handleDateQueryType(date string) string {
	if date == "All Time" {
		return ""
	} else if date == "Most Recent"{
		return ""
	} else {
		return buildDateQuery(date)
	}
}

// we want to store the dates as numerical values we can use in a range query.
// need to account for branching queries for all time and most recent, which should be diverted.
func buildDateQuery(date string) string {
	now := time.Now().UTC().UnixMilli()
	year := now - 31556952000 // this is technically not necessary, since databases will be sharded by year.
	month := now - 2628000000
	week := now - 604800000
	day := now - 86400000

	dateMap := map[string]int64{
		"Past Year": year,
		"Past Month": month,
		"Past Week": week,
		"Past Day": day,
	} 

	dateValue := dateMap[date]
	dateQuery := fmt.Sprintf("WHERE (w).date >%d", dateValue)
	return dateQuery
}

func buildWriterQuery(params url.Values) {

}
