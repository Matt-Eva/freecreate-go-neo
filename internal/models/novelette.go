package models

import (
	"errors"
	"fmt"
	"freecreate/internal/utils"
)

type Novelette struct {
	Writing
	Years []int
}

func (n Novelette) validateNovelette(year int)error{
	if len(n.Years) <= 0 || len(n.Years) > 1{
		err := "new novelette must only have original year within years property upon creation"
		return errors.New(err)
	}
	if n.Years[0] != year {
		err := "year added to novelette years upon creation does not match original novelette year"
		return errors.New(err)
	}
	if n.WritingType != "novelette"{
		err := fmt.Sprintf("writing type '%s' does not match novelette", n.WritingType)
		return errors.New(err)
	}
	return nil
}

func (n Novelette) newNoveletteParams()map[string]any{
	writingParams := utils.NeoParamsFromStruct(n)

	params := map[string]any{
		"writingParams": writingParams,
		"creatorId": n.CreatorId,
	}

	return params
}

type PostedNovelette struct {
	PostedWriting
}

func (p PostedNovelette) generateNovelette(year int)(Novelette, error){
	writing, err := p.generateWriting(year)
	if err != nil {
		return Novelette{}, err
	}

	years := []int{year}
	novelette := Novelette{
		Writing: writing,
		Years: years,
	}

	nErr := novelette.validateNovelette(year)
	if nErr != nil {
		return Novelette{}, nErr
	}

	return novelette, nil
}