package databases

import (
	"errors"
	"fmt"
	"freecreate/internal/utils"
)

func BuildSearchQuery(searchType, writingType, name, timeFrame string, genres, tags []string) (string, string, error) {
	if searchType == "writing" {
		return BuildWritingSearchQuery(writingType, name, timeFrame, genres, tags)
	} else if searchType == "writers" {
		return BuildWriterSearchQuery(name, genres, tags)
	} else {
		errorMessage := fmt.Sprintf("search type %s does not match valid search types", searchType)
		return "", "", errors.New(errorMessage)
	}
}

// WRITING SEARCH

func BuildWritingSearchQuery(writingType, name, timeFrame string, genres, tags []string) (string, string, error) {
	validatedWritingType, wErr := utils.ValidateWritingType(writingType)
	if wErr != nil {
		return "", "", wErr
	}

	validatedTimeFrame, tErr := utils.ValidateTimeFrame(timeFrame)
	if tErr != nil {
		return "", "", tErr
	}

	validatedGenres, gErr := utils.ValidateGenres(genres)
	if gErr != nil {
		return "", "", gErr
	}

	if name == "" && len(tags) == 0 && validatedTimeFrame != "mostRecent" {
		// search cache
		query := BuildRedisCacheQuery(validatedWritingType, validatedTimeFrame, validatedGenres)
		return query, "redis", nil
	} else if validatedTimeFrame == "mostRecent" {
		// query most recent database
		// order by date posted
		query := BuildMostRecentNeoQuery()
		return query, "neo", nil
	} else if validatedTimeFrame == "allTime" {
		// query all time database
		// order by absolute rank
		query := BuildAllTimeNeoQuery()
		return query, "neo", nil
	} else if utils.GetYearMap()[validatedTimeFrame] {
		// query specific year database
		// order by absolute rank and relative rank
		query := BuildSpecificYearNeoQuery()
		return query, "neo", nil
	} else if name != "" || len(tags) != 0 {
		// query most recent database
		// order by absolute rank and relative rank
		query := BuildStandardWritingNeoQuery(validatedWritingType, name, validatedTimeFrame, validatedGenres, tags)
		return query, "neo", nil
	} else {
		return "", "", errors.New("error")
	}
}

func BuildRedisCacheQuery(writingType, datePosted string, genres []string) string {
	return ""
}

func BuildMostRecentNeoQuery() string {
	return ""
}

func BuildAllTimeNeoQuery() string {
	return ""
}

func BuildSpecificYearNeoQuery() string {
	return ""
}

func BuildStandardWritingNeoQuery(writingType, name, timeFrame string, genres, tags []string) string {
	return ""
}

// WRITER SEARCH

func BuildWriterSearchQuery(name string, genres, tags []string) (string, string, error) {
	return "", "", nil
}
