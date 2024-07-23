package models

import (
	"fmt"
	"freecreate/internal/err"
)

type Universe struct {
	Writing
	Years []int
}

func (n Universe) validateUniverse(year int) err.Error {
	if len(n.Years) <= 0 || len(n.Years) > 1 {
		e := "new universe must only have original year within years property upon creation"
		return err.New(e)
	}
	if n.Years[0] != year {
		e := "year added to universe years upon creation does not match original universe year"
		return err.New(e)
	}
	if n.WritingType != "universe" {
		e := fmt.Sprintf("writing type '%s' does not match universe", n.WritingType)
		return err.New(e)
	}
	return err.Error{}
}

func MakeUniverse(p PostedWriting, year int) (Universe, err.Error) {
	writing, gErr := MakeWriting(p, year)
	if gErr.E != nil {
		return Universe{}, gErr
	}

	years := []int{year}
	universe := Universe{
		Writing: writing,
		Years:   years,
	}

	nErr := universe.validateUniverse(year)
	if nErr.E != nil {
		return Universe{}, nErr
	}

	return universe, err.Error{}
}
