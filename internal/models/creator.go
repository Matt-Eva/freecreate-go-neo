package models

import (
	"freecreate/internal/err"
	"freecreate/internal/utils"

	"github.com/google/uuid"
)

type Creator struct {
	Uid         string
	Name string
	CreatorId   string
	UserId      string
	ProfilePic  string
	About       string
	CreatedAt   int64
	UpdatedAt   int64
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

type NewCreator struct {
	Name string 
	CreatorId   string 
	About       string 
}

func GenerateCreator(userId string, n NewCreator) (Creator, err.Error) {
	var creator Creator
	utils.StructToStruct(n, creator)
	uid := uuid.New().String()
	creator.Uid = uid
	creator.UserId = userId
	creator.ProfilePic = ""

	cErr := creator.validateCreator()
	if cErr.E != nil {
		return Creator{}, cErr
	}

	return creator, err.Error{}
}

type UpdatedCreatorInfo struct {
	Uid string
	Name string
	CreatorId string
	About string
}

func (u UpdatedCreatorInfo) validateUpdatedCreator()err.Error {
	if u.CreatorId == ""{
		return err.New("creator id cannot be empty")
	}
	if u.Name == ""{
		return err.New("creator name cannot be empty")
	}
	if u.Uid == ""{
		return err.New("uid must be sent up with creator info")
	}

	return err.Error{}
}

type IncomingUpdatedCreatorInfo struct {
	Uid string
	Name string
	CreatorId string
	About string
}

func MakeUpdatedCreatorInfo(i IncomingUpdatedCreatorInfo)(UpdatedCreatorInfo, err.Error) {
	var u UpdatedCreatorInfo
	if e := utils.StructToStruct(i, u); e.E != nil {
		return u, e
	}
	if e := u.validateUpdatedCreator(); e.E != nil {
		return u, e
	}

	return u, err.Error{}
}


