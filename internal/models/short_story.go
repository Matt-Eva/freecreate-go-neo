package models

import (
	"fmt"
	"freecreate/internal/err"
)

type ShortStory struct {
	Writing
}

func (s ShortStory) validateNewShortStory() err.Error {
	if s.WritingType != "shortStory" {
		errorMsg := fmt.Sprintf("Writing type '%s' is not valid for a short Story; must be of type shortStory", s.WritingType)
		return err.New(errorMsg)
	}
	return err.Error{}
}



func MakeShortStory(year int) (ShortStory, err.Error) {

	shortStory := ShortStory{writing}

	

	return shortStory, err.Error{}
}
