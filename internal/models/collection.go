package models

import (
	"errors"
	"fmt"
)

type Collection struct {
	Writing
	Years []int
}

func (n Collection) validateCollection(year int) error {
	if len(n.Years) <= 0 || len(n.Years) > 1 {
		err := "new collection must only have original year within years property upon creation"
		return errors.New(err)
	}
	if n.Years[0] != year {
		err := "year added to collection years upon creation does not match original collection year"
		return errors.New(err)
	}
	if n.WritingType != "collection" {
		err := fmt.Sprintf("writing type '%s' does not match collection", n.WritingType)
		return errors.New(err)
	}
	return nil
}

func (n Collection) newCollectionParams() map[string]any {
	params := n.newWritingParams()
	return params
}

type PostedCollection struct {
	PostedWriting
}

func (p PostedCollection) generateCollection(year int) (Collection, error) {
	writing, err := p.generateWriting(year)
	if err != nil {
		return Collection{}, err
	}

	years := []int{year}
	collection := Collection{
		Writing: writing,
		Years:   years,
	}

	nErr := collection.validateCollection(year)
	if nErr != nil {
		return Collection{}, nErr
	}

	return collection, nil
}
