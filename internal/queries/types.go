package queries

type RetrievedWriting struct {
	Uid              string
	Title            string
	Description      string
	Genres           []string
	Tags             []string
	Author           string
	UniqueAuthorName string
	CreatorId        string
	Font             string
	WritingType      string
	Published        bool
}

type RetrievedChapterNoContent struct {
	Uid           string `bson:"uid"`
	Title         string `bson:"title"`
	WritingId     string `bson:"writing_id"`
	ChapterNumber int `bson:"chapter_number"`
}
