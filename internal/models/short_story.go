package models

import (
	"errors"
	"fmt"
)

type ShortStory struct {
	Writing
}

func (s ShortStory) validateNewShortStory(year int) error {
	if s.WritingType != "shortStory" {
		errorMsg := fmt.Sprintf("Writing type '%s' is not valid for a short Story; must be of type shortStory", s.WritingType)
		return errors.New(errorMsg)
	}
	return nil
}

func (s ShortStory) newShortStoryParams() map[string]any{
	writParams := s.newWritingParams()

	params := map[string]any{
		"writingParams": writParams,
		"creatorId": s.CreatorId,
	}

	return params
}

type PostedShortStory struct {
	PostedWriting
}

func (p PostedShortStory) generateShortStory(year int)(ShortStory, error){
	writing, err := p.generateWriting(year)
	if err != nil {
		return ShortStory{}, err
	}

	shortStory := ShortStory{writing}

	return shortStory, nil
}
