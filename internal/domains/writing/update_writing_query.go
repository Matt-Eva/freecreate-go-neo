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

func UpdateWritingQuery(ctx context.Context, neo neo4j.DriverWithContext, userId string, updateInfo UpdateWriting) (ReturnedWriting, int, err.Error) {
	status, aErr := auth.CheckAuthorizedUserWriting(ctx, neo, userId, updateInfo.Uid)
	if aErr.E != nil {
		return ReturnedWriting{}, status, aErr
	}
	status, uErr := auth.CheckAuthorizedUserCreator(ctx, neo, userId, updateInfo.CreatorId)
	if uErr.E != nil {
		return ReturnedWriting{}, status, uErr
	}

	query, qErr := buildUpdateWritingQuery(updateInfo.Genres, updateInfo.Tags)
	if qErr.E != nil {
		return ReturnedWriting{}, 500, qErr
	}

	params := utils.StructToMap(updateInfo)
	params["userId"] = userId

	db := os.Getenv("NEO_DB")
	if db == "" {
		return ReturnedWriting{}, 500, err.New("db env variable is empty")
	}

	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(db))
	if nErr != nil {
		return ReturnedWriting{}, 500, err.NewFromErr(nErr)
	}
	if len(result.Records) < 1 {
		return ReturnedWriting{}, 404, err.New("no records returned from database")
	}

	recordMap := result.Records[0].AsMap()
	var returnedWriting ReturnedWriting
	if e := utils.MapToStruct(recordMap, &returnedWriting); e.E != nil {
		return ReturnedWriting{}, 500, e
	}

	returnedWriting.Tags = updateInfo.Tags
	returnedWriting.Genres = updateInfo.Genres

	return returnedWriting, 200, err.Error{}
}

func buildUpdateWritingQuery(genres []string, tags []string) (string, err.Error) {
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

	genreLabels, gErr := utils.ValidateGenreLabels(genres)
	if gErr.E != nil {
		return "", gErr
	}

	hasTagLabel, hErr := utils.GetRelationshipLabel("HAS_TAG")
	if hErr.E != nil {
		return "", hErr
	}

	tagLabel, tErr := utils.GetNodeLabel("Tag")
	if tErr.E != nil {
		return "", tErr
	}

	writingNodeLabels := writingLabel
	for _, label := range genreLabels {
		writingNodeLabels += label
	}

	matchQuery := fmt.Sprintf("MATCH (u:%s {uid: $userId}) - [:%s] -> (:%s) - [cr:%s] -> (w:%s {uid: $uid})", userLabel, isCreatorLabel, creatorLabel, createdLabel, writingLabel)

	matchCreatorQuery := fmt.Sprintf("MATCH (u) - [:%s] -> (c:%s {uid: $creatorId})", isCreatorLabel, creatorLabel)

	optionalMatchTagQuery := fmt.Sprintf("OPTIONAL MATCH (w) - [h:%s] -> (:%s)", hasTagLabel, tagLabel)

	deleteHasTagRelQuery := "DELETE h "

	deleteCreatedRelQuery := "DELETE cr "

	createCreatedRelQuery := fmt.Sprintf("CREATE (c) -[:%s] -> (w)", createdLabel)

	setLabelsQuery := fmt.Sprintf("SET w:%s", writingNodeLabels)

	setQuery := `
		SET w.creatorId = $creatorId,
		w.title = $title,
		w.description = $description,
		w.font = $font,
		w.writingType = $writingType
	`

	tagQuery := fmt.Sprintf(`
		WITH w, c
		UNWIND $tags as tag
		MERGE (t:%s {tag: tag})
		MERGE (w) - [:%s] -> (t)
	`, tagLabel, hasTagLabel)

	returnQuery := `
		RETURN w.uid AS Uid,
		w.title AS Title,
		w.description AS Description,
		w.font AS Font,
		w.published AS Published,
		w.writingType AS WritingType,
		c.name AS Author,
		c.uniqueName AS UniqueAuthorName,
		c.uid As CreatorId
	`

	query := ""
	if len(tags) > 0 {
		query = matchQuery + matchCreatorQuery + optionalMatchTagQuery + deleteHasTagRelQuery + deleteCreatedRelQuery + createCreatedRelQuery + setLabelsQuery + setQuery + tagQuery + returnQuery
	} else {
		query = matchQuery + matchCreatorQuery + optionalMatchTagQuery + deleteHasTagRelQuery + deleteCreatedRelQuery + createCreatedRelQuery + setLabelsQuery + setQuery + returnQuery
	}

	return query, err.Error{}

}
