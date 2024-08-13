package queries

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"os"
	"runtime"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type RetrievedChapter struct {
	Title  string `bson:"title"`
	Number int    `bson:"chapter_number"`
	Uid    string `bson:"uid"`
}

type RetrievedNeoWriting struct {
	Uid         string
	Title       string
	Description string
	Genres      []string
	Tags        []string
	Author      string
	CreatorId   string
}

type RetrievedWriting struct {
	Uid         string
	Title       string
	Description string
	Author      string
	Genres      []string
	Tags        []string
	CreatorId   string
	Chapters    []RetrievedChapter
}

func GetWriting(ctx context.Context, neo neo4j.DriverWithContext, mongo *mongo.Client, creatorId, writingId string) (RetrievedWriting, err.Error) {
	neoChan := make(chan RetrievedNeoWriting)
	neoErrorChan := make(chan err.Error)
	go getNeoWriting(ctx, neo, creatorId, writingId, neoChan, neoErrorChan)

	mongoChan := make(chan []RetrievedChapter)
	mongoErrorChan := make(chan err.Error)
	go getMongoChapters(ctx, mongo, creatorId, writingId, mongoChan, mongoErrorChan)

	for e := range neoErrorChan {
		return RetrievedWriting{}, e
	}
	for e := range mongoErrorChan {
		return RetrievedWriting{}, e
	}

	if len(neoChan) > 1 {
		return RetrievedWriting{}, err.New("multiple writing nodes returned")
	}
	if len(mongoChan) > 1 {
		return RetrievedWriting{}, err.New("mutiple sets of chapters returned")
	}

	retrievedWriting := &RetrievedWriting{}
	for w := range neoChan {
		retrievedWriting.Uid = w.Uid
		retrievedWriting.Title = w.Title
		retrievedWriting.Description = w.Description
		retrievedWriting.Author = w.Author
		retrievedWriting.Genres = w.Genres
		retrievedWriting.Tags = w.Tags
		retrievedWriting.CreatorId = w.CreatorId
	}
	for c := range mongoChan {
		retrievedWriting.Chapters = c
	}

	return (*retrievedWriting), err.Error{}
}

func getNeoWriting(ctx context.Context, neo neo4j.DriverWithContext, creatorId, writingId string, neoChan chan RetrievedNeoWriting, errorChan chan err.Error) {
	defer close(neoChan)
	defer close(errorChan)

	neoQuery, qErr := buildNeoGetWritingQuery()
	if qErr.E != nil {
		errorChan <- qErr
		runtime.Goexit()
	}

	neoParams := map[string]any{
		"creatorId": creatorId,
		"writingId": writingId,
	}

	neoDb := os.Getenv("NEO_DB")
	if neoDb == "" {
		errorChan <- err.New("could not get neo db env variable")
		runtime.Goexit()
	}

	neoResult, nErr := neo4j.ExecuteQuery(ctx, neo, neoQuery, neoParams, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(neoDb))
	if nErr != nil {
		errorChan <- err.NewFromErr(nErr)
		runtime.Goexit()
	}
	if len(neoResult.Records) < 1 {
		errorChan <- err.New("no records returned from database query")
		runtime.Goexit()
	}

	resultMap := make(map[string]any)
	tagSlice := make([]string, 0)

	for _, record := range neoResult.Records {
		recordMap := record.AsMap()
		for key, val := range recordMap {
			if key == "Tag" {
				stringVal, ok := val.(string)
				if !ok {
					errorChan <- err.New("tag field from record could not be converted to string")
					runtime.Goexit()
				}

				tagSlice = append(tagSlice, stringVal)
			}

			_, ok := resultMap[key]
			if !ok {
				resultMap[key] = val
			}
		}
	}

	resultMap["Tags"] = tagSlice
	var retrievedNeoWriting RetrievedNeoWriting
	if e := utils.MapToStruct(resultMap, &retrievedNeoWriting); e.E != nil {
		errorChan <- e
		runtime.Goexit()
	}

	neoChan <- retrievedNeoWriting
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
	matchCreatorQuery := fmt.Sprintf("MATCH (w) <- [:%s] - (c:%s)", createdLabel, creatorLabel)
	matchTagQuery := fmt.Sprintf("MATCH (w) - [:%s] -> (t:%s)", hasTagRelationship, tagLabel)

	returnQuery := `
		RETURN w.uid AS Uid,
		w.title AS Title,
		w.description AS Description,
		labels(w) AS Genres,
		t.tag AS Tag,
		c.name AS Author,
		c.uid As CreatorId
	`

	query := matchWritQuery + matchCreatorQuery + matchTagQuery + returnQuery

	return query, err.Error{}

}

func getMongoChapters(ctx context.Context, mongo *mongo.Client, creatorId, writingId string, mongoChan chan []RetrievedChapter, errorChan chan err.Error) {
	defer close(mongoChan)
	defer close(errorChan)

	filter := bson.D{{"creator_id", creatorId}, {"neo_id", writingId}}

	chapterColl := mongo.Database("freecreate").Collection("chapters")

	queryOptions := options.Find().SetProjection(bson.D{{"title", 1}, {"chatper_number", 1}})

	cursor, mErr := chapterColl.Find(ctx, filter, queryOptions)
	if mErr != nil {
		e := err.NewFromErr(mErr)
		errorChan <- e
		runtime.Goexit()
	}

	var results []RetrievedChapter
	if e := cursor.All(ctx, &results); e != nil {
		newE := err.NewFromErr(e)
		errorChan <- newE
		runtime.Goexit()
	}

	mongoChan <- results

}
