package models

import (
	"freecreate/internal/err"
	"freecreate/internal/utils"
)

type User struct {
	Uid         string
	DisplayName string
	Username    string
	Email       string
	Password    string
	ProfilePic  string
	BirthYear   int
	BirthMonth  int
	BirthDay    int
}

func (u User) ValidateUser() err.Error {
	if u.Uid == "" {
		e := "user uid is empty"
		return err.New(e)
	}
	if u.DisplayName == "" {
		e := "user display name is empty"
		return err.New(e)
	}
	if u.Username == "" {
		e := "user username is empty"
		return err.New(e)
	}
	if u.Email == "" {
		e := "user email is empty"
		return err.New(e)
	}
	if u.Password == "" {
		e := "user password is empty"
		return err.New(e)
	}
	if u.ProfilePic != "" {
		e := "profile pic must be empty - not currently accepting images"
		return err.New(e)
	}
	if u.BirthYear == 0 || u.BirthYear < 1900 {
		e := "birth year error"
		return err.New(e)
	}
	if u.BirthMonth == 0 || u.BirthMonth > 12 || u.BirthMonth < 0 {
		e := "birth month error"
		return err.New(e)
	}
	if u.BirthDay == 0 || u.BirthDay > 31 || u.BirthMonth < 0{
		e := "birth day error"
		return err.New(e)
	}
	return err.Error{}
}

func (u User) NewUserParams() map[string]any {
	userParams := utils.NeoParamsFromStruct(u)

	return map[string]any{"userParams": userParams}
}

type PostedUser struct {
	DisplayName          string `json:"displayName"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	ProfilePic           string `json:"profilePic"`
	BirthYear            string `json:"birthYear"`
	BirthMonth           string `json:"birthMonth"`
	BirthDay             string `json:"birthDay"`
}

// func (p PostedUser) GenerateUser() (User, err.Error) {
// 	if p.Password != p.PasswordConfirmation {
// 		e := "password and password confirmation do not match"
// 		return User{}, err.New(e)
// 	}

// 	newUser := User{
// 		Uid:         uuid.New().String(),
// 		DisplayName: p.DisplayName,
// 		Username:    p.Username,
// 		Email:       p.Email,
// 		Password:    p.Password,
// 		ProfilePic:  p.ProfilePic,
// 		BirthYear:   p.BirthYear,
// 		BirthMonth:  p.BirthMonth,
// 		BirthDay:    p.BirthDay,
// 	}

// 	vErr := newUser.validateUser()
// 	if vErr.E != nil {
// 		return newUser, vErr
// 	}

// 	return newUser, err.Error{}
// }
