package models

import (
	"errors"
	"freecreate/internal/utils"

	"github.com/google/uuid"
)


type Tag struct {
	Uid string
	Tag string
}

func (t Tag) validateTag() error {
	if t.Uid == ""{
		err := "tag uid cannot be empty"
		return errors.New(err)
	}
	if t.Tag == ""{
		err := "tag name cannot be empty"
		return errors.New(err)
	}
	return nil
}

func (t Tag) newTagParams() map[string]any{
	params := utils.NeoParamsFromStruct(t)
	return params
}

type PostedTag struct {
	Tag string
}

func (p PostedTag) generateTag () Tag{
	return Tag{
		Uid: uuid.New().String(),
		Tag: p.Tag,
	}
}