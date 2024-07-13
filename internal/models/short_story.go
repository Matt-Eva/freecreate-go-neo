package models

import (
	"errors"
	"fmt"
)

type ShortStory struct {
	Writing
}

func (s ShortStory) validateNewShortStory(year int) error {
	err := s.validateNewWriting(year)
	if err != nil {
		return err
	}

	if s.WritingType != "shortStory" {
		errorMsg := fmt.Sprintf("Writing type '%s' is not valid for a short Story; must be of type shortStory", s.WritingType)
		return errors.New(errorMsg)
	}
	return nil
}
