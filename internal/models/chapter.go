package models

import (
	"freecreate/internal/err"
	"time"

	"github.com/google/uuid"
)

// chapters exist only in Mongo, not in Neo4j
// shard key is creatorid
// query for chapter is creatorid/neoid/uid

type Chapter struct {
	Uid           string         `bson:"uid"`
	WritingId     string         `bson:"writing_id"`
	Title         string         `bson:"title"`
	ChapterNumber int            `bson:"chapter_number"`
	Published     bool           `bson:"published"`
	CreatedAt     int64          `bson:"created_at"`
	UpdatedAt     int64          `bson:"updated_at`
	Content       map[string]any `bson:"content"`
}

func (c Chapter) validateChapter() err.Error {
	if c.Uid == "" {
		return err.New("uid cannot be empty")
	}
	if c.WritingId == "" {
		return err.New("WritingId cannot be empty")
	}
	if c.Title == "" {
		return err.New("Title cannot be empty")
	}
	if c.ChapterNumber == 0 {
		return err.New("ChapterNumber cannot be empty")
	}
	if c.Published {
		return err.New("chapter cannot be published on creation")
	}
	if c.CreatedAt == 0 {
		return err.New("CreatedAt cannot be empty")
	}
	if c.UpdatedAt == 0 {
		return err.New("UpdatedAt cannot be empty")
	}
	if len(c.Content) > 0 {
		return err.New("content must start be empty")
	}

	return err.Error{}
}

type PostedChapter struct {
	WritingId     string
	Title         string
	ChapterNumber int
}

func MakeChapter(p PostedChapter) (Chapter, err.Error) {
	chapter := &Chapter{}

	chapter.WritingId = p.WritingId
	chapter.Title = p.Title
	chapter.ChapterNumber = p.ChapterNumber

	now := time.Now().UnixMilli()
	chapter.CreatedAt = now
	chapter.UpdatedAt = now

	uid := uuid.New().String()
	chapter.Uid = uid

	vErr := chapter.validateChapter()
	if vErr.E != nil {
		return *chapter, vErr
	}

	return *chapter, err.Error{}
}
