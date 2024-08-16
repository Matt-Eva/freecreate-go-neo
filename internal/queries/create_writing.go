package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/utils"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type CreatedWriting struct {
	Uid              string
	CreatorId        string
	Title            string
	Description      string
	Font             string
	Author           string
	UniqueAuthorName string
	Genres           []string
	Tags             []string
}

func CreateWriting(ctx context.Context, neo neo4j.DriverWithContext, userId string, writing models.Writing, genres, tags []string) (CreatedWriting, int, err.Error) {
	status, aErr := checkAuthorizedUserCreator(ctx, neo, userId, writing.CreatorId)
	if aErr.E != nil {
		return CreatedWriting{}, status, aErr
	}

	query, qErr := buildCreateWritingQuery(genres, tags)
	if qErr.E != nil {
		return CreatedWriting{}, 500, qErr
	}

	writingParams := utils.StructToMap(writing)

	params := map[string]any{
		"userId":        userId,
		"creatorId":     writing.CreatorId,
		"tags":          tags,
		"writingParams": writingParams,
	}

	db := os.Getenv("NEO_DB")
	if db == "" {
		return CreatedWriting{}, 500, err.New("db env variable is empty")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return CreatedWriting{}, 500, err.NewFromErr(nErr)
	}
	if len(result.Records) < 1 {
		return CreatedWriting{}, 500, err.New("db query returned no records")
	}

	recordMap := result.Records[0].AsMap()
	var createdWriting CreatedWriting
	if e := utils.MapToStruct(recordMap, &createdWriting); e.E != nil {
		return CreatedWriting{}, 500, e
	}

	createdWriting.Tags = tags
	createdWriting.Genres = genres

	return createdWriting, 201, err.Error{}
}

func buildCreateWritingQuery(genres, tags []string) (string, err.Error) {
	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	isCreatorLabel, iErr := GetRelationshipLabel("IS_CREATOR")
	if iErr.E != nil {
		return "", iErr
	}

	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	createdLabel, crErr := GetRelationshipLabel("CREATED")
	if crErr.E != nil {
		return "", crErr
	}

	writingLabel, wErr := GetNodeLabel("Writing")
	if wErr.E != nil {
		return "", wErr
	}

	genreLabels, gErr := buildGenreLabels(genres)
	if gErr.E != nil {
		return "", gErr
	}

	tagLabel, tErr := GetNodeLabel("Tag")
	if tErr.E != nil {
		return "", tErr
	}

	hasTagLabel, htErr := GetRelationshipLabel("HAS_TAG")
	if htErr.E != nil {
		return "", htErr
	}

	writingLabels := writingLabel + genreLabels

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId}) - [:%s] -> (c:%s {uid: $creatorId})", userLabel, isCreatorLabel, creatorLabel)
	createQuery := fmt.Sprintf("CREATE (c) - [:%s] -> (w:%s $writingParams)", createdLabel, writingLabels)

	tagQuery := fmt.Sprintf(`
	WITH w, c
	UNWIND $tags as tag
	MERGE (t:%s {tag: tag})
	MERGE (w) - [:%s] -> (t)
	`, tagLabel, hasTagLabel)

	returnQuery := `
		RETURN w.uid AS Uid, 
		w.creatorId AS CreatorId, 
		w.title AS Title, 
		w.description AS Description, 
		w.font AS Font, 
		c.name AS Author, 
		c.uniqueName AS UniqueAuthorName
	`

	query := ""
	if len(tags) > 0{
		query = matchQuery + createQuery + tagQuery + returnQuery
	} else {
		query = matchQuery + createQuery + returnQuery
	}

	return query, err.Error{}
}
