package models

import (
	"errors"
	"fmt"
)

type Universe struct {
	Writing
	Years []int
}

func (n Universe) validateUniverse(year int) error {
	if len(n.Years) <= 0 || len(n.Years) > 1 {
		err := "new universe must only have original year within years property upon creation"
		return errors.New(err)
	}
	if n.Years[0] != year {
		err := "year added to universe years upon creation does not match original universe year"
		return errors.New(err)
	}
	if n.WritingType != "universe" {
		err := fmt.Sprintf("writing type '%s' does not match universe", n.WritingType)
		return errors.New(err)
	}
	return nil
}

func (n Universe) newUniverseParams() map[string]any {
	params := n.newWritingParams()

	return params
}

type PostedUniverse struct {
	PostedWriting
}

func (p PostedUniverse) generateuniverse(year int) (Universe, error) {
	writing, err := p.generateWriting(year)
	if err != nil {
		return Universe{}, err
	}

	years := []int{year}
	universe := Universe{
		Writing: writing,
		Years:   years,
	}

	nErr := universe.validateUniverse(year)
	if nErr != nil {
		return Universe{}, nErr
	}

	return universe, nil
}
