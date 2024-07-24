package models

import (
	"errors"

	"github.com/google/uuid"
)

type Tag struct {
	Uid string
	Tag string
}

func (t Tag) validateTag() error {
	if t.Uid == "" {
		err := "tag uid cannot be empty"
		return errors.New(err)
	}
	if t.Tag == "" {
		err := "tag name cannot be empty"
		return errors.New(err)
	}
	return nil
}

type PostedTag struct {
	Tag string
}

func (p PostedTag) generateTag() Tag {
	return Tag{
		Uid: uuid.New().String(),
		Tag: p.Tag,
	}
}
