package writing_handler

import "freecreate/internal/err"

type ReturnedWriting struct {
	Uid              string   `json:"uid"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	Author           string   `json:"author"`
	Font             string   `json:"font"`
	UniqueAuthorName string   `json:"uniqueAuthorName"`
	Genres           []string `json:"genres"`
	Tags             []string `json:"tags"`
	CreatorId        string   `json:"creatorId"`
	Published 	bool `json:"published"`
}

func validateReturnedWriting(r ReturnedWriting, genres, tags []string) err.Error {
	if r.Uid == "" {
		return err.New("returned writing Uid cannot be empty")
	}
	if r.Title == "" {
		return err.New("return writing title cannot be empty")
	}
	if r.Author == "" {
		return err.New("return writing Author cannot be empty")
	}
	if r.Font == "" {
		return err.New("return writing Font cannot be empty")
	}
	if r.UniqueAuthorName == "" {
		return err.New("return writing unique author name cannot be empty")
	}
	if r.CreatorId == "" {
		return err.New("return writing CreatorId cannot be empty")
	}
	if len(r.Genres) != len(genres) {
		return err.New("return writing genre number does not match number of input genres")
	}
	if len(r.Tags) != len(tags) {
		return err.New("return writing Tag number does not match number of input Tags")
	}

	return err.Error{}
}
