package models

import (
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Uid         string
	DisplayName string
	Username    string
	Email       string
	Password    string
	ProfilePic  string
	Birthday    int64
}

func (u User) validateUser() err.Error {
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
	if u.Birthday == 0 {
		e := "birthday cannot be empty"
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

func (p PostedUser) GenerateUser() (User, err.Error) {
	if p.Password != p.PasswordConfirmation {
		e := "password and password confirmation do not match"
		return User{}, err.New(e)
	}

	year, yErr := strconv.Atoi(p.BirthYear)
	if yErr != nil {
		e := "could not convert birth year to number"
		return User{}, err.New(e)
	}

	month, mErr := strconv.Atoi(p.BirthMonth)
	if mErr != nil {
		e := "could not convert birth month to number"
		return User{}, err.New(e)
	}

	day, dErr := strconv.Atoi(p.BirthDay)
	if dErr != nil {
		e := "could not convert birth day to number"
		return User{}, err.New(e)
	}

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	birthday := date.UnixMilli()

	newUser := User{
		Uid:         uuid.New().String(),
		DisplayName: p.DisplayName,
		Username:    p.Username,
		Email:       p.Email,
		Password:    p.Password,
		ProfilePic:  p.ProfilePic,
		Birthday:    birthday,
	}

	vErr := newUser.validateUser()
	if vErr.E != nil {
		return newUser, vErr
	}

	return newUser, err.Error{}
}
