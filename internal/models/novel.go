package models

import (
	"fmt"
	"freecreate/internal/err"
)

type Novel struct {
	Writing
	Years []int
}

func (n Novel) validateNovel(year int) err.Error {
	if len(n.Years) <= 0 || len(n.Years) > 1 {
		e := "new novel must only have original year within years property upon creation"
		return err.New(e)
	}
	if n.Years[0] != year {
		e := "year added to novel years upon creation does not match original novel year"
		return err.New(e)
	}
	if n.WritingType != "novel" {
		e := fmt.Sprintf("writing type '%s' does not match novel", n.WritingType)
		return err.New(e)
	}

	return err.Error{}
}

func (n Novel) newNovelParams() map[string]any {
	params := n.newWritingParams()

	return params
}

type PostedNovel struct {
	PostedWriting
}

func (p PostedNovel) generateNovel(year int) (Novel, err.Error) {
	writing, gErr := p.generateWriting(year)
	if gErr.E != nil {
		return Novel{}, gErr
	}

	years := []int{year}
	novel := Novel{
		Writing: writing,
		Years:   years,
	}

	nErr := novel.validateNovel(year)
	if nErr.E != nil {
		return Novel{}, nErr
	}

	return novel, err.Error{}
}
