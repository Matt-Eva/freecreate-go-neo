package models

import (
	"freecreate/internal/err"
	"freecreate/internal/utils"

	"github.com/google/uuid"
)

type Creator struct {
	Uid         string
	CreatorName string
	CreatorId   string
	UserId      string
	ProfilePic  string
	About       string
}

func (c Creator) validateCreator() err.Error {
	if c.Uid == "" {
		return err.New("creator uid cannot be empty")
	}
	if c.CreatorName == "" {
		return err.New("creator name cannot be empty")
	}
	if c.CreatorId == "" {
		return err.New("creator unique identifier - user provided - cannot be empty")
	}
	if c.UserId == "" {
		return err.New("creator userId cannot be empty")
	}
	if c.ProfilePic != "" {
		return err.New("creator profile must be empty - not currently accepting profile pics")
	}

	return err.Error{}
}

type NewCreator struct {
	CreatorName string 
	CreatorId   string 
	About       string 
}

func GenerateCreator(userId string, n NewCreator) (Creator, err.Error) {
	var creator Creator
	utils.StructToStruct(n, creator)
	uid := uuid.New().String()
	creator.Uid = uid

	cErr := creator.validateCreator()
	if cErr.E != nil {
		return Creator{}, cErr
	}

	return creator, err.Error{}
}

