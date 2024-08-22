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
