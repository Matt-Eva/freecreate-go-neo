package models

import (
	"errors"
	"fmt"
)

type Novel struct {
	Writing
	Years []int
}

func (n Novel) validateNovel(year int)error{
	if len(n.Years) <= 0 || len(n.Years) > 1{
		err := "new novel must only have original year within years property upon creation"
		return errors.New(err)
	}
	if n.Years[0] != year {
		err := "year added to novel years upon creation does not match original novel year"
		return errors.New(err)
	}
	if n.WritingType != "novel"{
		err := fmt.Sprintf("writing type '%s' does not match novel", n.WritingType)
		return errors.New(err)
	}
	return nil
}

func (n Novel) newNovelParams()map[string]any{
	params := n.newWritingParams()


	return params
}

type PostedNovel struct {
	PostedWriting
}

func (p PostedNovel) generateNovel(year int)(Novel, error){
	writing, err := p.generateWriting(year)
	if err != nil {
		return Novel{}, err
	}

	years := []int{year}
	novel := Novel{
		Writing: writing,
		Years: years,
	}

	nErr := novel.validateNovel(year)
	if nErr != nil {
		return Novel{}, nErr
	}

	return novel, nil
}
