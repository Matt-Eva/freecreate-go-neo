package writing

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"
	"slices"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetUserWriting(ctx context.Context, neo neo4j.DriverWithContext, userId string) ([]ReturnedWriting, int, err.Error) {
	query, qErr := buildGetUserWritingQuery()
	if qErr.E != nil {
		return []ReturnedWriting{}, 500, qErr
	}

	params := map[string]any{
		"userId": userId,
	}

	db := os.Getenv("NEO_DB")
	if db == "" {
		return []ReturnedWriting{}, 500, err.New("could not get db environment variable")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return []ReturnedWriting{}, 500, err.NewFromErr(nErr)
	}

	writingHash := make(map[string]*ReturnedWriting)

	for _, record := range result.Records {
		recordMap := record.AsMap()
		resultMap := make(map[string]any)
		uid, ok := recordMap["Uid"]
		if !ok {
			return []ReturnedWriting{}, 500, err.New("record does not have Uid attribute")
		}

		uidString, ok := uid.(string)
		if !ok {
			return []ReturnedWriting{}, 500, err.New("could not convert uid to string")
		}

		_, exists := writingHash[uidString]
		if !exists {
			writingHash[uidString] = &ReturnedWriting{}
		}
		retWrit := writingHash[uidString]

		for key, val := range recordMap {
			if key == "Genres" {
				genres := make([]string, 0)
				if slice, ok := val.([]any); ok {
					for _, g := range slice {
						if strG, ok := g.(string); ok {
							genres = append(genres, strG)
						}
					}
				}
				retWrit.Genres = genres
				continue
			}

			if key == "Tag" {
				if tag, ok := val.(string); ok {
					retWrit.Tags = append(retWrit.Tags, tag)
				}
				continue
			}

			_, ok := resultMap[key]
			if !ok {
				resultMap[key] = val
			}
		}

		if e := utils.MapToStruct(resultMap, retWrit); e.E != nil {
			return []ReturnedWriting{}, 500, e
		}
	}

	writing := make([]ReturnedWriting, 0)

	for _, val := range writingHash {
		writing = append(writing, (*val))
	}

	slices.SortFunc(writing, func(i, j ReturnedWriting) int {
		return strings.Compare(i.Title, j.Title)
	})

	return writing, 200, err.Error{}
}

func buildGetUserWritingQuery() (string, err.Error) {
	userLabel, uErr := utils.GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	isCreatorLabel, iErr := utils.GetRelationshipLabel("IS_CREATOR")
	if iErr.E != nil {
		return "", iErr
	}

	creatorLabel, cErr := utils.GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	createdLabel, crErr := utils.GetRelationshipLabel("CREATED")
	if crErr.E != nil {
		return "", crErr
	}

	writingLabel, wErr := utils.GetNodeLabel("Writing")
	if wErr.E != nil {
		return "", wErr
	}

	hasTagLabel, hErr := utils.GetRelationshipLabel("HAS_TAG")
	if hErr.E != nil {
		return "", hErr
	}

	tagLabel, tErr := utils.GetNodeLabel("Tag")
	if tErr.E != nil {
		return "", tErr
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId}) - [:%s] -> (c:%s) - [:%s] -> (w:%s)", userLabel, isCreatorLabel, creatorLabel, createdLabel, writingLabel)
	tagMatchQuery := fmt.Sprintf("OPTIONAL MATCH (w) - [:%s] -> (t:%s)", hasTagLabel, tagLabel)
	returnQuery := `
		RETURN w.uid AS Uid,
		w.title AS Title,
		w.description AS Description,
		w.font AS Font,
		w.published AS Published,
		labels(w) AS Genres,
		c.name AS Author,
		c.uniqueName AS UniqueAuthorName,
		c.uid AS CreatorId,
		t.tag AS Tag
	`

	query := matchQuery + tagMatchQuery + returnQuery

	return query, err.Error{}
}
