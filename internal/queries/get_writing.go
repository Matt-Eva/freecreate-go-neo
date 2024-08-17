package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)



type RetrievedNeoWriting struct {
	Uid         string
	Title       string
	Description string
	Genres      []string
	Tags        []string
	Author      string
	UniqueAuthorName string
	CreatorId   string
	Font string
}



func GetWriting(ctx context.Context, neo neo4j.DriverWithContext, creatorId, writingId string) (RetrievedNeoWriting, int, err.Error) {

	retrievedNeoWriting,status, nErr := getNeoWriting(ctx, neo, creatorId, writingId)
	if nErr.E != nil {
		return retrievedNeoWriting, status, nErr
	}

	return retrievedNeoWriting, status, err.Error{}
}

func getNeoWriting(ctx context.Context, neo neo4j.DriverWithContext, creatorId, writingId string) (RetrievedNeoWriting, int, err.Error) {


	neoQuery, qErr := buildNeoGetWritingQuery()
	if qErr.E != nil {
		return RetrievedNeoWriting{}, 500, qErr
	}

	neoParams := map[string]any{
		"creatorId": creatorId,
		"writingId": writingId,
	}

	neoDb := os.Getenv("NEO_DB")
	if neoDb == "" {
		return RetrievedNeoWriting{}, 500, err.New("could not get neo db env variable")
	}

	neoResult, nErr := neo4j.ExecuteQuery(ctx, neo, neoQuery, neoParams, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(neoDb))
	if nErr != nil {
		return RetrievedNeoWriting{}, 500, err.NewFromErr(nErr)
	}
	if len(neoResult.Records) < 1 {
		return RetrievedNeoWriting{}, 404, err.New("no records returned from database")
	}

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
					return RetrievedNeoWriting{}, 500, err.New("tag field from record could not be converted to string")
				}

				fmt.Println(val)

				tagSlice = append(tagSlice, stringVal)
				continue
			} else if key == "Genres"{
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

	
	var retrievedNeoWriting RetrievedNeoWriting
	if e := utils.MapToStruct(resultMap, &retrievedNeoWriting); e.E != nil {
		return retrievedNeoWriting, 500, e
	}

	labels := make([]string, 0)

	for val, _ := range genreMap{
		labels = append(labels, val)
	}

	genres, _ := validateGenreLabels(labels)

	retrievedNeoWriting.Tags = tagSlice
	retrievedNeoWriting.Genres = genres

	return retrievedNeoWriting, 200, err.Error{}
}

func buildNeoGetWritingQuery() (string, err.Error) {
	writingLabel, wErr := GetNodeLabel("Writing")
	if wErr.E != nil {
		return "", wErr
	}

	createdLabel, rErr := GetRelationshipLabel("CREATED")
	if rErr.E != nil {
		return "", rErr
	}

	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	hasTagRelationship, hErr := GetRelationshipLabel("HAS_TAG")
	if hErr.E != nil {
		return "", hErr
	}

	tagLabel, tErr := GetNodeLabel("Tag")
	if tErr.E != nil {
		return "", tErr
	}

	matchWritQuery := fmt.Sprintf("MATCH (w:%s {uid: $writingId})", writingLabel)
	matchCreatorQuery := fmt.Sprintf("MATCH (w) <- [:%s] - (c:%s {uid: $creatorId})", createdLabel, creatorLabel)
	matchTagQuery := fmt.Sprintf("OPTIONAL MATCH (w) - [:%s] -> (t:%s)", hasTagRelationship, tagLabel)

	returnQuery := `
		RETURN w.uid AS Uid,
		w.title AS Title,
		w.description AS Description,
		w.font AS Font,
		labels(w) AS Genres,
		t.tag AS Tag,
		c.name AS Author,
		c.uniqueName AS UniqueAuthorName,
		c.uid As CreatorId
	`

	query := matchWritQuery + matchCreatorQuery + matchTagQuery + returnQuery

	return query, err.Error{}

}
