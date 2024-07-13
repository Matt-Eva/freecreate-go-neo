package models

import (
	"errors"
	"fmt"
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
	Libraries    int64
	Likes        int64
	Views        int64
	Donations    int64
	Rank         int64
	RelRank      int64
}

func (w Writing) validateNewWriting() error{
	if w.Uid == ""{
		e := fmt.Sprintf("uid cannot be empty")
		return errors.New(e)
	}
	if w.Title == ""{
		e := fmt.Sprintf("title cannot be empty")
		return errors.New(e)
	}
	if w.Description == ""{
		e := fmt.Sprintf("description cannot be empty")
		return errors.New(e)
	}

	return nil
}

func (w Writing) newWritingParams() map[string]any{
	writingParams := map[string]any {
		"uid": w.Uid,
	}

	paramMap := map[string]any {
		"writingParams": writingParams,
		"creatorId": w.CreatorId,
	}

	return paramMap
}

type PostedWriting struct {
	Title string
	Description string
	WritingType string
	Thumbnail string
	CreatorId string
}

func (p PostedWriting) generateWriting() Writing {
	now := time.Now().UnixMilli()
	newWriting := Writing {
		Uid: uuid.New().String(),
		Title: p.Title,
		Description: p.Description,
		Thumbnail: "",
		WritingType: p.WritingType,
		CreatorId: p.CreatorId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return newWriting
}

type UpdateWritingInfo struct {
	Uid [16]byte
	Title string
	Description string
	Thumbnail string
	WritingType string
}



type ShortStory struct {
	Writing
}

func (s ShortStory) validateNewShortStory() error {
	err := s.validateNewWriting()
	if err != nil {
		return err
	}

	if s.WritingType != "shortStory"{
		errorMsg := fmt.Sprintf("Writing type '%s' is not valid for a short Story; must be of type shortStory", s.WritingType)
		return errors.New(errorMsg)
	}
	return nil
}

type Novel struct {
	Writing
}
