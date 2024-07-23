package models

import (
	"freecreate/internal/err"
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
	BirthYear   int
	BirthMonth  int
	BirthDay    int
	CreatedAt int64
	UpdatedAt int64
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
	if u.CreatedAt == 0 {
		return err.New("created at cannot be 0")
	}
	if u.UpdatedAt == 0 {
		return err.New("updated at cannot be 0")
	}
	return err.Error{}
}

type PostedUser struct {
	DisplayName string `json:"displayName"`
	Username	string `json:"username"`
	Email string `json:"email"`
	BirthDay int `json:"birthday"`
	BirthYear int `json:"birthYear"`
	BirthMonth int `json:"birthMonth"`
	ProfilePic string `json:"profilePic"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func MakeNewUser(p PostedUser)(User, err.Error){
	if p.Password != p.PasswordConfirmation{
		e := err.New("password and password confirmation do not match")
		return User{}, e
	}
	uid := uuid.New().String()
	now := time.Now().UnixMilli()
	u := User {
		DisplayName: p.DisplayName,
		Password: p.Password,
		Username: p.Username,
		BirthYear: p.BirthYear,
		BirthMonth: p.BirthMonth,
		BirthDay: p.BirthDay,
		Uid: uid,
		Email: p.Email,
		ProfilePic: p.ProfilePic,
		CreatedAt: now,
		UpdatedAt: now,
	}

	vErr := u.validateUser()
	if vErr.E != nil {
		return User{}, vErr
	} 

	return u, err.Error{}
}