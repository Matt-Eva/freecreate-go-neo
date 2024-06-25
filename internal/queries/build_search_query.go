package databases

import (
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

func BuildWritingSearchQuery(writingType, name, datePosted string, genres, tags []string){
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
	} else if datePosted == "Most Recent"{
		// query most recent database
		// order by date posted
		BuildMostRecentNeoQuery()
	} else if datePosted == "All Time"{
		// query all time database
		// order by absolute rank
		BuildAllTimeNeoQuery()
	} else if utils.GetYearMap()[datePosted] {
		// query specific year database
		// order by absolute rank and relative rank
		BuildSpecificYearNeoQuery()
	} else if (name != "" || len(tags) != 0) && checkDateMap[datePosted] {
		// query most recent database
		// order by absolute rank and relative rank
		BuildStandardWritingNeoQuery()
	} else {
		// return error
	}
}

func BuildRedisCacheQuery(){

}

func BuildMostRecentNeoQuery(){

}

func BuildAllTimeNeoQuery(){

}

func BuildSpecificYearNeoQuery(){

}

func BuildStandardWritingNeoQuery(){

}

func BuildWriterSearchQuery(name string, genres, tags []string){

}