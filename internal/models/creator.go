package models

import "github.com/google/uuid"

type Creator struct {
	Uid string
	Name string
	UserId string
	ProfilePic string
	About string
}

type PostedCreator struct {
	Name string
	ProfilePic string
	About string
}

func (p PostedCreator) generateCreator(userId string)(Creator){
	creator := Creator {
		Uid: uuid.New().String(),
		Name: p.Name,
		ProfilePic: p.ProfilePic,
		About: p.About,
		UserId: userId,
	}
	
	return creator
}