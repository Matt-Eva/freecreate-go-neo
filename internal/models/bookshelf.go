package models

import (
	"errors"

	"github.com/google/uuid"
)

type Bookshelf struct {
	Uid string
	Name string
	UserId string
}

func (b Bookshelf) validateNewBookshelf() error{
	if b.Uid == ""{
		e := "bookshelf uid cannot be nil"
		return errors.New(e)
	}
	if b.Name == ""{
		e := "bookshelf name cannot be nil"
		return errors.New(e)
	}
	if b.UserId == ""{
		e := "bookshelf UserId cannot be nil"
		return errors.New(e)
	}
	
	return nil
}

type PostedBookshelf struct {
	Name string `json:"name"`
}

func (p PostedBookshelf) generateBookshelf(userId string)(Bookshelf, error){
	bookshelf := Bookshelf{
		Uid: uuid.New().String(),
		Name: p.Name,
		UserId: userId,
	}

	vErr := bookshelf.validateNewBookshelf()
	if vErr != nil {
		return Bookshelf{}, vErr
	}
	
	return bookshelf, nil
}