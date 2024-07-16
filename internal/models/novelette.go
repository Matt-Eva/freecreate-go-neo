package models

import (
	"fmt"
	"freecreate/internal/err"
)

type Novelette struct {
	Writing
	Years []int
}

func (n Novelette) validateNovelette(year int) err.Error {
	if len(n.Years) <= 0 || len(n.Years) > 1 {
		e := "new novelette must only have original year within years property upon creation"
		return err.New(e)
	}
	if n.Years[0] != year {
		e := "year added to novelette years upon creation does not match original novelette year"
		return err.New(e)
	}
	if n.WritingType != "novelette" {
		e := fmt.Sprintf("writing type '%s' does not match novelette", n.WritingType)
		return err.New(e)
	}
	return err.Error{}
}

func (n Novelette) newNoveletteParams() map[string]any {
	params := n.newWritingParams()

	return params
}

type PostedNovelette struct {
	PostedWriting
}

func (p PostedNovelette) generateNovelette(year int) (Novelette, err.Error) {
	writing, gErr := p.generateWriting(year)
	if gErr.E != nil {
		return Novelette{}, gErr
	}

	years := []int{year}
	novelette := Novelette{
		Writing: writing,
		Years:   years,
	}

	nErr := novelette.validateNovelette(year)
	if nErr.E != nil {
		return Novelette{}, nErr
	}

	return novelette, err.Error{}
}
