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

func (s ShortStory) NewShortStoryParams() map[string]any {
	params := s.newWritingParams()

	return params
}

type PostedShortStory struct {
	PostedWriting
}

func (p PostedShortStory) GenerateShortStory(year int) (ShortStory, err.Error) {
	writing, gErr := p.generateWriting(year)
	if gErr.E != nil {
		return ShortStory{}, gErr
	}

	shortStory := ShortStory{writing}

	vErr := shortStory.validateNewShortStory()
	if vErr.E != nil {
		return ShortStory{}, vErr
	}

	return shortStory, err.Error{}
}
