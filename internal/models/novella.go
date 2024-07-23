package models

import (
	"fmt"
	"freecreate/internal/err"
)

type Novella struct {
	Writing
	Years []int
}

func (n Novella) validateNovella(year int) err.Error {
	if len(n.Years) <= 0 || len(n.Years) > 1 {
		e := "new novella must only have original year within years property upon creation"
		return err.New(e)
	}
	if n.Years[0] != year {
		e := "year added to novella years upon creation does not match original novella year"
		return err.New(e)
	}
	if n.WritingType != "novella" {
		e := fmt.Sprintf("writing type '%s' does not match novella", n.WritingType)
		return err.New(e)
	}
	return err.Error{}
}

func MakeNovella(p PostedWriting, year int) (Novella, err.Error) {
	writing, gErr := MakeWriting(p, year)
	if gErr.E != nil {
		return Novella{}, gErr
	}

	years := []int{year}
	novella := Novella{
		Writing: writing,
		Years:   years,
	}

	nErr := novella.validateNovella(year)
	if nErr.E != nil {
		return Novella{}, nErr
	}

	return novella, err.Error{}
}
