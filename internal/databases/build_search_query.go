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

	} else if datePosted == "All Time"{

	} else if utils.GetYearMap()[datePosted] {

	} else if name != "" || len(tags) != 0 {

	} else {
		// return error
	}
}

func BuildWriterSearchQuery(name string, genres, tags []string){

}