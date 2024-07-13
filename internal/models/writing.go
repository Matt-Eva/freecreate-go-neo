package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Writing struct {
	Uid          string
	Title        string
	Description  string
	Thumbnail    string
	WritingType  string
	CreatorId    string
	CreatedAt    int64
	UpdatedAt    int64
	LibraryCount int64
	Likes        int64
	Views        int64
	Donations    int64
	Rank         int64
	RelRank      int64
	OriginalYear int
}

func (w Writing) validateNewWriting(year int) error {
	if w.Uid == "" {
		e := "uid cannot be empty"
		return errors.New(e)
	}
	if w.Title == "" {
		e := "title cannot be empty"
		return errors.New(e)
	}
	if w.Description == "" {
		e := "description cannot be empty"
		return errors.New(e)
	}
	if w.Thumbnail != "" {
		e := "thumbnail must be empty - not accepting thumbnail images at present"
		return errors.New(e)
	}
	if w.WritingType == "" {
		e := "writing type cannot be empty"
		return errors.New(e)
	}
	if w.CreatorId == "" {
		e := "creator id cannot be emtpy"
		return errors.New(e)
	}
	if w.CreatedAt == 0 {
		e := "server side error - no value inserted for created at"
		return errors.New(e)
	}
	if w.UpdatedAt == 0 {
		e := "server side error - no value inserted for updated at"
		return errors.New(e)
	}
	if w.UpdatedAt != w.CreatedAt {
		e := "server side error - created at and updated at must match"
		return errors.New(e)
	}
	if w.LibraryCount != 0 {
		e := "library count cannot be greater than 0 for new writing"
		return errors.New(e)
	}
	if w.Likes != 0 {
		e := "like count cannot be greater than 0 for new writing"
		return errors.New(e)
	}
	if w.Donations != 0 {
		e := "donation count cannot be greater than 0 for new writing"
		return errors.New(e)
	}
	if w.Views != 0 {
		e := "view count cannot be greater than 0 for new writing"
		return errors.New(e)
	}
	if w.Rank != 0 {
		e := "rank count cannot be greater than 0 for new writing"
		return errors.New(e)
	}
	if w.RelRank != 0 {
		e := "relrank count cannot be greater than 0 for new writing"
		return errors.New(e)
	}
	if w.OriginalYear != year || w.OriginalYear == 0 {
		e := "server side error - Original year does not match current year or is empty"
		return errors.New(e)
	}

	return nil
}

func (w Writing) newWritingParams() (map[string]any, error) {
	v := reflect.ValueOf(w)
	t := v.Type()
	writingParams := make(map[string]any)

	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i).Name
		mapField := strings.ToLower(structField[:1]) + structField[1:]
		writingParams[mapField] = v.Field(i).Interface()
	}

	paramMap := map[string]any{
		"writingParams": writingParams,
		"creatorId":     w.CreatorId,
	}

	return paramMap, nil
}

type PostedWriting struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	WritingType string `json:"writingType"`
	Thumbnail   string `json:"thumbnail"`
	CreatorId   string `json:"creatorId"`
}

func (p PostedWriting) generateWriting(year int) Writing {
	now := time.Now().UnixMilli()
	newWriting := Writing{
		Uid:          uuid.New().String(),
		Title:        p.Title,
		Description:  p.Description,
		Thumbnail:    "",
		WritingType:  p.WritingType,
		CreatorId:    p.CreatorId,
		CreatedAt:    now,
		UpdatedAt:    now,
		OriginalYear: year,
	}

	return newWriting
}

type UpdateWritingInfo struct {
	Uid         [16]byte
	Title       string
	Description string
	Thumbnail   string
	WritingType string
}
