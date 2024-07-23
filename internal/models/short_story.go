package models

import (
	"fmt"
	"freecreate/internal/err"
)

type ShortStory struct {
	Writing
}

func (s ShortStory) validateShortStory() err.Error {
	if s.WritingType != "shortStory" {
		errorMsg := fmt.Sprintf("Writing type '%s' is not valid for a short Story; must be of type shortStory", s.WritingType)
		return err.New(errorMsg)
	}
	
	return err.Error{}
}

func MakeShortStory(p PostedWriting, year int) (ShortStory, err.Error) {

	w, wErr := MakeWriting(p, year)
	if wErr.E != nil {
		return ShortStory{}, wErr
	}

	shortStory := ShortStory{w}

	vErr := shortStory.validateShortStory()
	if vErr.E != nil {
		return ShortStory{}, vErr
	}

	return shortStory, err.Error{}
}
