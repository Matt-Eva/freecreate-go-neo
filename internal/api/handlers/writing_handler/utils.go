package writing_handler

import (
	"freecreate/internal/err"
	"freecreate/internal/queries"
)

func convertRetrievwedWritingToReturnedWriting(retrieved queries.RetrievedWriting)(ReturnedWriting, err.Error){
	r := &ReturnedWriting{}

	r.Uid = retrieved.Uid
	r.Author = retrieved.Author
	r.Title = retrieved.Title
	r.Description = retrieved.Description
	r.Font = retrieved.Font
	r.UniqueAuthorName = retrieved.UniqueAuthorName
	r.Genres = retrieved.Genres
	r.Tags = retrieved.Tags
	r.Published = retrieved.Published
	r.CreatorId = retrieved.CreatorId

	if e:= validateReturnedWriting(*r, r.Genres, r.Tags); e.E != nil {
		return *r, e
	}

	return *r, err.Error{}
}