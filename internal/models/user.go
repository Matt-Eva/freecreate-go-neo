package models

import (
	"errors"
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

func (u User) validateUser() error {
	if u.Uid == "" {
		err := "user uid is empty"
		return errors.New(err)
	}
	if u.DisplayName == "" {
		err := "user display name is empty"
		return errors.New(err)
	}
	if u.Username == "" {
		err := "user username is empty"
		return errors.New(err)
	}
	if u.Email == "" {
		err := "user email is empty"
		return errors.New(err)
	}
	if u.Password == "" {
		err := "user password is empty"
		return errors.New(err)
	}
	if u.ProfilePic != "" {
		err := "profile pic must be empty - not currently accepting images"
		return errors.New(err)
	}
	if u.Birthday == 0 {
		err := "birthday cannot be empty"
		return errors.New(err)
	}
	return nil
}

func (u User) newUserParams() map[string]any {
	userParams := utils.NeoParamsFromStruct(u)

	return userParams
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

func (p PostedUser) generateUser() (User, error) {
	if p.Password != p.PasswordConfirmation {
		err := "password and password confirmation do not match"
		return User{}, errors.New(err)
	}

	year, yErr := strconv.Atoi(p.BirthYear)
	if yErr != nil {
		err := "could not convert birth year to number"
		return User{}, errors.New(err)
	}

	month, mErr := strconv.Atoi(p.BirthMonth)
	if mErr != nil {
		err := "could not convert birth month to number"
		return User{}, errors.New(err)
	}

	day, dErr := strconv.Atoi(p.BirthDay)
	if dErr != nil {
		err := "could not convert birth day to number"
		return User{}, errors.New(err)
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
	if vErr != nil {
		return newUser, vErr
	}

	return newUser, nil
}
