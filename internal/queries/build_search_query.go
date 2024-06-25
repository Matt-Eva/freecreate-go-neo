package databases

import (
	"errors"
	"freecreate/internal/utils"
)

func BuildSearchQuery(searchType, writingType, name, datePosted string, genres, tags []string){
	if searchType == "writing"{
		BuildWritingSearchQuery(writingType, name, datePosted, genres, tags)
	} else if searchType == "writers"{
		BuildWriterSearchQuery(name, genres, tags)
	} else {
		// return error
	}
}

func BuildWritingSearchQuery(writingType, name, datePosted string, genres, tags []string)(string, string, error){
	validatedType := utils.ValidateWritingType(writingType)
	validatedGenres := utils.ValidateGenres(genres)

	checkDateMap := map[string]bool {
		"Past Day": true,
		"Past Week": true,
		"Past Month": true,
		"Past Year": true,
		"All Time": true,
	}

	if name == "" && len(tags) == 0 && checkDateMap[datePosted] {
		// search cache
		query := BuildRedisCacheQuery()
		return query, "redis", nil
	} else if datePosted == "Most Recent"{
		// query most recent database
		// order by date posted
		query := BuildMostRecentNeoQuery()
		return query, "neo", nil
	} else if datePosted == "All Time"{
		// query all time database
		// order by absolute rank
		query := BuildAllTimeNeoQuery()
		return query, "neo", nil
	} else if utils.GetYearMap()[datePosted] {
		// query specific year database
		// order by absolute rank and relative rank
		query := BuildSpecificYearNeoQuery()
		return query, "neo", nil
	} else if (name != "" || len(tags) != 0) && checkDateMap[datePosted] {
		// query most recent database
		// order by absolute rank and relative rank
		query := BuildStandardWritingNeoQuery()
		return query, "neo", nil
	} else {
		return "", "", errors.New("error")
	}
}

func BuildRedisCacheQuery()string{
return ""
}

func BuildMostRecentNeoQuery()string{
return ""
}

func BuildAllTimeNeoQuery()string{
return ""
}

func BuildSpecificYearNeoQuery()string{
return ""
}

func BuildStandardWritingNeoQuery()string{
return ""
}

func BuildWriterSearchQuery(name string, genres, tags []string){

}