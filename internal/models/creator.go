package models

import (
	"freecreate/internal/err"

	"github.com/google/uuid"
)

type Creator struct {
	Uid        string
	CreatorName       string
	CreatorId  string
	UserId     string
	ProfilePic string
	About      string
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

type PostedCreator struct {
	CreatorName       string `json:"name"`
	CreatorId  string `json:"creatorId"`
	About      string `json:"about"`
}

func GenerateCreator(userId string, p PostedCreator) (Creator, err.Error) {
	creator := Creator{
		Uid:        uuid.New().String(),
		CreatorName:       p.CreatorName,
		CreatorId:  p.CreatorId,
		About:      p.About,
		UserId:     userId,
	}

	cErr := creator.validateCreator()
	if cErr.E != nil {
		return Creator{}, cErr
	}

	return creator, err.Error{}
}
