package writing

import (
	"context"
	"fmt"
	"freecreate/internal/domains/auth"
	"freecreate/internal/err"
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
	WritingType      string
	Published        bool
}

func CreateWritingQuery(ctx context.Context, neo neo4j.DriverWithContext, userId string, writing Writing, genres, tags []string) (ReturnedWriting, int, err.Error) {
	status, aErr := auth.CheckAuthorizedUserCreator(ctx, neo, userId, writing.CreatorId)
	if aErr.E != nil {
		return ReturnedWriting{}, status, aErr
	}

	query, qErr := buildCreateWritingQuery(genres, tags)
	if qErr.E != nil {
		return ReturnedWriting{}, 500, qErr
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
		return ReturnedWriting{}, 500, err.New("db env variable is empty")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return ReturnedWriting{}, 500, err.NewFromErr(nErr)
	}
	if len(result.Records) < 1 {
		return ReturnedWriting{}, 500, err.New("db query returned no records")
	}

	recordMap := result.Records[0].AsMap()
	var createdWriting ReturnedWriting
	if e := utils.MapToStruct(recordMap, &createdWriting); e.E != nil {
		return ReturnedWriting{}, 500, e
	}

	createdWriting.Tags = tags
	createdWriting.Genres = genres

	return createdWriting, 201, err.Error{}
}

func buildCreateWritingQuery(genres, tags []string) (string, err.Error) {
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

	genreLabels, gErr := utils.BuildGenreLabels(genres)
	if gErr.E != nil {
		return "", gErr
	}

	tagLabel, tErr := utils.GetNodeLabel("Tag")
	if tErr.E != nil {
		return "", tErr
	}

	hasTagLabel, htErr := utils.GetRelationshipLabel("HAS_TAG")
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
		w.published AS Published,
		w.writingType AS WritingType,
		c.name AS Author, 
		c.uniqueName AS UniqueAuthorName
	`

	query := ""
	if len(tags) > 0 {
		query = matchQuery + createQuery + tagQuery + returnQuery
	} else {
		query = matchQuery + createQuery + returnQuery
	}

	return query, err.Error{}
}
