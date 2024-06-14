package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func SearchHandler(w http.ResponseWriter, r *http.Request, neo, mongo, redis string)  {
	params := r.URL.Query()
	detectSearchType(params)
}

func detectSearchType(params url.Values) ([]string, error) {
	if params["searchType"][0] == "writing"{
		 buildWritingQuery(params)
	} else if params["searchType"][0] == "writer" {
		buildWriterQuery(params)
	}
	return []string{}, errors.New("")
}

func buildWritingQuery(params url.Values) ([]string, error){
	// date := params["date"]
	genres := params["genre"]
	tags := params["tag"]
	title := params["writingTitle"]
	// writingType := params["writingType"]
	if len(tags) == 0 && len(title) == 0 {
		fmt.Println("hit cache")
	} else {
		genreQuery := buildWritingGenreQuery(genres)
		fmt.Println(genreQuery)
	}

	return []string{}, nil
}

func buildWritingGenreQuery (genres []string) string{
	genreLabels := []string{"Action", "Adventure", "Comedy", "Drama", "HistoricalFiction", "Horror", "Fantasy", "LiteraryFiction", "MagicalRealism", "Mystery", "Realism", "Romance", "SliceOfLife", "SocialFiction", "Superhero", "Thriller"}

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



func buildWriterQuery(params url.Values){

}