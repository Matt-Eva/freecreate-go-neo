package models

import (
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"time"

	"github.com/google/uuid"
)

// ====== Base User Model ======

type User struct {
	Uid        string
	UniqueName string
	Username   string
	Email      string
	Password   string
	ProfilePic string
	BirthYear  int
	BirthMonth int
	BirthDay   int
	CreatedAt  int64
	UpdatedAt  int64
}

func (u User) validateUser() err.Error {
	if u.Uid == "" {
		e := "user uid is empty"
		return err.New(e)
	}
	if u.UniqueName == "" {
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
	if u.BirthDay == 0 || u.BirthDay > 31 || u.BirthMonth < 0 {
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
	UniqueName           string
	Username             string
	Email                string
	BirthDay             int
	BirthYear            int
	BirthMonth           int
	ProfilePic           string
	Password             string
	PasswordConfirmation string
}

func GenerateUser(p PostedUser) (User, err.Error) {
	if p.Password != p.PasswordConfirmation {
		e := err.New("password and password confirmation do not match")
		return User{}, e
	}
	uid := uuid.New().String()
	now := time.Now().UnixMilli()
	u := User{
		UniqueName: p.UniqueName,
		Password:   p.Password,
		Username:   p.Username,
		BirthYear:  p.BirthYear,
		BirthMonth: p.BirthMonth,
		BirthDay:   p.BirthDay,
		Uid:        uid,
		Email:      p.Email,
		ProfilePic: p.ProfilePic,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	vErr := u.validateUser()
	if vErr.E != nil {
		return User{}, vErr
	}

	return u, err.Error{}
}

// ====== Updated User Info ======

type UpdatedUserInfo struct {
	UserId     string
	Username   string
	Email      string
	BirthDay   int
	BirthYear  int
	BirthMonth int
}

func (u UpdatedUserInfo) validateUpdateUserInfo() err.Error {
	if u.UserId == "" {
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
	if u.BirthYear == 0 || u.BirthYear < 1900 {
		e := "birth year error"
		return err.New(e)
	}
	if u.BirthMonth == 0 || u.BirthMonth > 12 || u.BirthMonth < 0 {
		e := "birth month error"
		return err.New(e)
	}
	if u.BirthDay == 0 || u.BirthDay > 31 || u.BirthMonth < 0 {
		e := "birth day error"
		return err.New(e)
	}
	return err.Error{}
}

type PatchedUser struct {
	UserId     string
	Username   string
	Email      string
	BirthDay   int
	BirthYear  int
	BirthMonth int
}

func GenerateUpdatedUserInfo(p PatchedUser) (UpdatedUserInfo, err.Error) {
	var updatedUser UpdatedUserInfo
	if e := utils.StructToStruct(p, &updatedUser); e.E != nil {
		return updatedUser, e
	}

	vErr := updatedUser.validateUpdateUserInfo()
	if vErr.E != nil {
		return updatedUser, vErr
	}

	return updatedUser, err.Error{}
}
