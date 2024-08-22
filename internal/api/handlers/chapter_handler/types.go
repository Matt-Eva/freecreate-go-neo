package chapter_handler

type ReturnChapterNoContent struct {
	Uid           string `json:"uid"`
	Title         string `json:"title"`
	WritingId     string `json:"writingId"`
	ChapterNumber int    `json:"chapterNumber"`
}
