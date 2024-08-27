package writing

import (
	"freecreate/internal/err"
	"time"

	"github.com/google/uuid"
)

var titleLength = 200
var descriptionLength = 1000

type Writing struct {
	Uid          string
	Title        string
	Description  string
	Thumbnail    string
	WritingType  string
	CreatorId    string
	Font         string
	CreatedAt    int64
	UpdatedAt    int64
	LibraryCount int64
	Likes        int64
	Views        int64
	Donations    int64
	Rank         int64
	RelRank      int64
	OriginalYear int
	Years        []int
	Published    bool
}

func (w Writing) validateNewWriting(year int) err.Error {
	if w.Uid == "" {
		e := "uid cannot be empty"
		return err.New(e)
	}
	if w.Font == "" {
		return err.New("font cannot be empty")
	}
	if w.Title == "" {
		e := "title cannot be empty"
		return err.New(e)
	}
	if len(w.Title) > titleLength {
		e := "title exceeds character limit"
		return err.New(e)
	}
	if len(w.Description) > descriptionLength {
		e := "description exceeds character limit"
		return err.New(e)
	}
	if w.Thumbnail != "" {
		e := "thumbnail must be empty - not accepting thumbnail images at present"
		return err.New(e)
	}
	if w.WritingType == "" {
		e := "writing type cannot be empty"
		return err.New(e)
	}
	if w.CreatorId == "" {
		e := "creator id cannot be empty"
		return err.New(e)
	}
	if w.CreatedAt == 0 {
		e := "server side error - no value inserted for created at"
		return err.New(e)
	}
	if w.UpdatedAt == 0 {
		e := "server side error - no value inserted for updated at"
		return err.New(e)
	}
	if w.UpdatedAt != w.CreatedAt {
		e := "server side error - created at and updated at must match"
		return err.New(e)
	}
	if w.LibraryCount != 0 {
		e := "library count cannot be greater than 0 for new writing"
		return err.New(e)
	}
	if w.Likes != 0 {
		e := "like count cannot be greater than 0 for new writing"
		return err.New(e)
	}
	if w.Donations != 0 {
		e := "donation count cannot be greater than 0 for new writing"
		return err.New(e)
	}
	if w.Views != 0 {
		e := "view count cannot be greater than 0 for new writing"
		return err.New(e)
	}
	if w.Rank != 0 {
		e := "rank count cannot be greater than 0 for new writing"
		return err.New(e)
	}
	if w.RelRank != 0 {
		e := "relrank count cannot be greater than 0 for new writing"
		return err.New(e)
	}
	if w.OriginalYear != year {
		e := "server side error - Original year does not match current year or is empty"
		return err.New(e)
	}
	if w.Published {
		e := "writing cannot be set to published upon its creation"
		return err.New(e)
	}

	return err.Error{}
}

func MakeWriting(p PostedWriting, year int) (Writing, err.Error) {
	now := time.Now().UnixMilli()
	newWriting := Writing{
		Uid:          uuid.New().String(),
		Title:        p.Title,
		Description:  p.Description,
		Thumbnail:    "",
		WritingType:  p.WritingType,
		CreatorId:    p.CreatorId,
		Font:         p.Font,
		CreatedAt:    now,
		UpdatedAt:    now,
		OriginalYear: year,
	}

	vErr := newWriting.validateNewWriting(year)
	if vErr.E != nil {
		return Writing{}, vErr
	}

	return newWriting, err.Error{}
}

type UpdateWriting struct {
	Uid         string
	CreatorId   string
	Title       string
	Description string
	Genres      []string
	Tags        []string
	Font        string
	WritingType string
}

func (u UpdateWriting) validateUpdateWriting() err.Error {
	if u.Uid == "" {
		return err.New("uid cannot be empty")
	}
	if u.CreatorId == "" {
		return err.New("CreatorId cannot be empty")
	}
	if u.Title == "" {
		return err.New("Title cannot be empty")
	}
	if u.Font == "" {
		return err.New("Font cannot be empty")
	}
	if u.WritingType == "" {
		return err.New("Writing type cannot be empty")
	}
	return err.Error{}
}

type PatchedWriting struct {
	Uid         string
	CreatorId   string
	Title       string
	Description string
	Genres      []string
	Tags        []string
	Font        string
	WritingType string
}

func MakeUpdateWriting(p PatchedWriting) (UpdateWriting, err.Error) {
	u := &UpdateWriting{}

	u.Uid = p.Uid
	u.CreatorId = p.CreatorId
	u.Title = p.Title
	u.Description = p.Description
	u.Genres = p.Genres
	u.Tags = p.Tags
	u.Font = p.Font
	u.WritingType = p.WritingType

	vErr := u.validateUpdateWriting()
	if vErr.E != nil {
		return *u, vErr
	}

	return *u, err.Error{}
}
