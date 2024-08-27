package chapters

type ReturnChapterNoContent struct {
	Uid           string `json:"uid" bson:"uid"`
	Title         string `json:"title" bson:"title"`
	WritingId     string `json:"writingId bson:"writing_id""`
	ChapterNumber int    `json:"chapterNumber" bson:"chapter_number"`
}
