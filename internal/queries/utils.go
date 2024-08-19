package queries

import (
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func convertNeoMapToRetrievedWriting(neoResult neo4j.EagerResult)(RetrievedWriting, int, err.Error){
	resultMap := make(map[string]any)
	tagSlice := make([]string, 0)
	genreMap := make(map[string]bool, 0)

	for _, record := range neoResult.Records {
		recordMap := record.AsMap()
		for key, val := range recordMap {
			if val == nil {
				continue
			}

			if key == "Tag" {
				stringVal, ok := val.(string)
				if !ok {
					return RetrievedWriting{}, 500, err.New("tag field from record could not be converted to string")
				}

				fmt.Println(val)

				tagSlice = append(tagSlice, stringVal)
				continue
			} else if key == "Genres" {
				if genres, ok := val.([]any); ok {
					for _, genre := range genres {
						if g, ok := genre.(string); ok {
							genreMap[g] = true
						}
					}
				}
				continue
			}

			_, ok := resultMap[key]
			if !ok {
				resultMap[key] = val
			}
		}
	}

	var retrievedNeoWriting RetrievedWriting
	if e := utils.MapToStruct(resultMap, &retrievedNeoWriting); e.E != nil {
		return retrievedNeoWriting, 500, e
	}

	labels := make([]string, 0)

	for key := range genreMap {
		labels = append(labels, key)
	}

	genres, _ := validateGenreLabels(labels)

	retrievedNeoWriting.Tags = tagSlice
	retrievedNeoWriting.Genres = genres

	return retrievedNeoWriting, 200, err.Error{}
}