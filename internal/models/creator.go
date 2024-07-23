package models

import (
	"freecreate/internal/err"
	"freecreate/internal/utils"

	"github.com/google/uuid"
)

type Creator struct {
	Uid        string
	Name       string
	CreatorId  string
	UserId     string
	ProfilePic string
	About      string
}

func (c Creator) validateCreator() err.Error {
	if c.Uid == "" {
		return err.New("creator uid cannot be empty")
	}
	if c.Name == "" {
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

func (c Creator) NewCreatorParams() map[string]any {
	params := utils.NeoParamsFromStruct(c)
	return map[string]any{"creatorParams": params, "userId": params["userId"]}
}

type PostedCreator struct {
	Name       string `json:"name"`
	CreatorId  string `json:"creatorId"`
	ProfilePic string `json:"profilePic"`
	About      string `json:"about"`
}

func (p PostedCreator) GenerateCreator(userId string) (Creator, err.Error) {
	creator := Creator{
		Uid:        uuid.New().String(),
		Name:       p.Name,
		CreatorId:  p.CreatorId,
		ProfilePic: p.ProfilePic,
		About:      p.About,
		UserId:     userId,
	}

	cErr := creator.validateCreator()
	if cErr.E != nil {
		return Creator{}, cErr
	}

	return creator, err.Error{}
}

