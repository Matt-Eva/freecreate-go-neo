package models

import (
	"errors"
	"fmt"
)

type Novella struct {
	Writing
	Years []int
}

func (n Novella) validateNovella(year int) error {
	if len(n.Years) <= 0 || len(n.Years) > 1 {
		err := "new novella must only have original year within years property upon creation"
		return errors.New(err)
	}
	if n.Years[0] != year {
		err := "year added to novella years upon creation does not match original novella year"
		return errors.New(err)
	}
	if n.WritingType != "novella" {
		err := fmt.Sprintf("writing type '%s' does not match novella", n.WritingType)
		return errors.New(err)
	}
	return nil
}

func (n Novella) newNovellaParams() map[string]any {
	params := n.newWritingParams()

	return params
}

type PostedNovella struct {
	PostedWriting
}

func (p PostedNovella) generateNovella(year int) (Novella, error) {
	writing, err := p.generateWriting(year)
	if err != nil {
		return Novella{}, err
	}

	years := []int{year}
	novella := Novella{
		Writing: writing,
		Years:   years,
	}

	nErr := novella.validateNovella(year)
	if nErr != nil {
		return Novella{}, nErr
	}

	return novella, nil
}
