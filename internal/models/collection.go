package models

import (
	"fmt"
	"freecreate/internal/err"
)

type Collection struct {
	Writing
	Years []int
}

func (n Collection) validateCollection(year int) err.Error {
	if len(n.Years) <= 0 || len(n.Years) > 1 {
		e := "new collection must only have original year within years property upon creation"
		return err.New(e)
	}
	if n.Years[0] != year {
		e := "year added to collection years upon creation does not match original collection year"
		return err.New(e)
	}
	if n.WritingType != "collection" {
		e := fmt.Sprintf("writing type '%s' does not match collection", n.WritingType)
		return err.New(e)
	}
	return err.Error{}
}

func (n Collection) newCollectionParams() map[string]any {
	params := n.newWritingParams()
	return params
}

type PostedCollection struct {
	PostedWriting
}

func (p PostedCollection) generateCollection(year int) (Collection, err.Error) {
	writing, gErr := p.generateWriting(year)
	if gErr.E != nil {
		return Collection{}, gErr
	}

	years := []int{year}
	collection := Collection{
		Writing: writing,
		Years:   years,
	}

	nErr := collection.validateCollection(year)
	if nErr.E != nil {
		return Collection{}, nErr
	}

	return collection, err.Error{}
}
