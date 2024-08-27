package writing

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetWritingQuery(ctx context.Context, neo neo4j.DriverWithContext, writingId string) (ReturnedWriting, int, err.Error) {

	retrievedNeoWriting, status, nErr := getNeoWriting(ctx, neo, writingId)
	if nErr.E != nil {
		return retrievedNeoWriting, status, nErr
	}

	return retrievedNeoWriting, status, err.Error{}
}

func getNeoWriting(ctx context.Context, neo neo4j.DriverWithContext, writingId string) (ReturnedWriting, int, err.Error) {

	neoQuery, qErr := buildNeoGetWritingQuery()
	if qErr.E != nil {
		return ReturnedWriting{}, 500, qErr
	}

	neoParams := map[string]any{
		"writingId": writingId,
	}

	neoDb := os.Getenv("NEO_DB")
	if neoDb == "" {
		return ReturnedWriting{}, 500, err.New("could not get neo db env variable")
	}

	neoResult, nErr := neo4j.ExecuteQuery(ctx, neo, neoQuery, neoParams, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(neoDb))
	if nErr != nil {
		return ReturnedWriting{}, 500, err.NewFromErr(nErr)
	}
	if len(neoResult.Records) < 1 {
		return ReturnedWriting{}, 404, err.New("no records returned from database")
	}

	// return convertNeoMapToReturnedWriting(*neoResult)

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
					return ReturnedWriting{}, 500, err.New("tag field from record could not be converted to string")
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

	var retrievedNeoWriting ReturnedWriting
	if e := utils.MapToStruct(resultMap, &retrievedNeoWriting); e.E != nil {
		return retrievedNeoWriting, 500, e
	}

	labels := make([]string, 0)

	for key := range genreMap {
		labels = append(labels, key)
	}

	genres, _ := utils.ValidateGenreLabels(labels)

	retrievedNeoWriting.Tags = tagSlice
	retrievedNeoWriting.Genres = genres

	return retrievedNeoWriting, 200, err.Error{}
}

func buildNeoGetWritingQuery() (string, err.Error) {
	writingLabel, wErr := utils.GetNodeLabel("Writing")
	if wErr.E != nil {
		return "", wErr
	}

	createdLabel, rErr := utils.GetRelationshipLabel("CREATED")
	if rErr.E != nil {
		return "", rErr
	}

	creatorLabel, cErr := utils.GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	hasTagRelationship, hErr := utils.GetRelationshipLabel("HAS_TAG")
	if hErr.E != nil {
		return "", hErr
	}

	tagLabel, tErr := utils.GetNodeLabel("Tag")
	if tErr.E != nil {
		return "", tErr
	}

	matchWritQuery := fmt.Sprintf("MATCH (w:%s {uid: $writingId})", writingLabel)
	matchCreatorQuery := fmt.Sprintf("MATCH (w) <- [:%s] - (c:%s)", createdLabel, creatorLabel)
	matchTagQuery := fmt.Sprintf("OPTIONAL MATCH (w) - [:%s] -> (t:%s)", hasTagRelationship, tagLabel)

	returnQuery := `
		RETURN w.uid AS Uid,
		w.title AS Title,
		w.description AS Description,
		w.font AS Font,
		w.published AS Published,
		w.writingType AS WritingType,
		labels(w) AS Genres,
		t.tag AS Tag,
		c.name AS Author,
		c.uniqueName AS UniqueAuthorName,
		c.uid As CreatorId
	`

	query := matchWritQuery + matchCreatorQuery + matchTagQuery + returnQuery

	return query, err.Error{}

}
