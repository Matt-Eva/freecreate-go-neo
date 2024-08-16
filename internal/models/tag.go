package models

import (
	"freecreate/internal/err"
	"freecreate/internal/utils"
)

type Tag struct {
	Tag string
}

func (t Tag) validateTag() err.Error {
	if t.Tag == "" {
		e := "tag name cannot be empty"
		return err.New(e)
	}
	return err.Error{}
}

type PostedTag struct {
	Tag string
}

func (p PostedTag) GenerateTag() (Tag, err.Error) {
	var tag Tag
	if tErr := utils.StructToStruct(p, &tag); tErr.E != nil {
		return tag, tErr
	}

	if e := tag.validateTag(); e.E != nil {
		return tag, e
	}
	
	return tag, err.Error{}
}
