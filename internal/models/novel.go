package models

type Novel struct {
	Writing
	Years []int
}

type PostedNovel struct {
	PostedWriting
}
