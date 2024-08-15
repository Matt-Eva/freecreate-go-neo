package middleware

import (
	"fmt"
	"freecreate/internal/err"
	"net/http"
	"os"

	"github.com/rbcervilla/redisstore/v9"
)

type AuthenticatedUser struct {
	Uid        string `json:"uid"`
	UniqueName     string `json:"uniqueName"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	BirthDay   int    `json:"birthDay"`
	BirthYear  int    `json:"birthYear"`
	BirthMonth int    `json:"birthMonth"`
	ProfilePic string `json:"profilePic"`
}

func (a AuthenticatedUser) validateAuthenticatedUser() err.Error {
	if a.Uid == "" {
		return err.New("authenticated user Uid cannot be empty")
	}
	if a.UniqueName == "" {
		return err.New("authenticated user UniqueName cannot be empty")
	}
	if a.Username == "" {
		return err.New("authenticated user Username cannot be empty")
	}
	if a.Email == "" {
		return err.New("authenticated user Email cannot be empty")
	}
	if a.BirthDay == 0 {
		return err.New("authenticated user BirthDay cannot be empty")
	}
	if a.BirthMonth == 0 {
		return err.New("authenticated user BirthMonth cannot be empty")
	}
	if a.BirthYear == 0 {
		return err.New("authenticated user BirthYear cannot be empty")
	}

	return err.Error{}
}

func AuthenticateUser(r *http.Request, store *redisstore.RedisStore) (AuthenticatedUser, err.Error) {
	sessionName := os.Getenv("USER_SESSION")
	userSession, gErr := store.Get(r, sessionName)
	if gErr != nil {
		return AuthenticatedUser{}, err.NewFromErr(gErr)
	}

	uid := userSession.Values["uid"]
	uniqueName := userSession.Values["uniqueName"]
	username := userSession.Values["username"]
	email := userSession.Values["email"]
	birthDay := userSession.Values["birthDay"]
	birthMonth := userSession.Values["birthMonth"]
	birthYear := userSession.Values["birthYear"]
	profilePic := userSession.Values["profilePic"]

	if username == nil || uniqueName == nil || uid == nil || email == nil || birthDay == nil || birthMonth == nil || birthYear == nil || profilePic == nil {
		msg := fmt.Sprintf(
			"user session attribute(s) nil\n: username: %T\n uniqueName: %T\n uid: %T\n email: %T\n birthDay: %T\n birthMonth: %T\n birthYear: %T\n profilePic: %T",
			username, uniqueName, uid, email, birthDay, birthMonth, birthYear, profilePic)
		return AuthenticatedUser{}, err.New(msg)
	}

	usernameS, ok := username.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session username to string")
	}

	uniqueNameS, ok := uniqueName.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session uniqueName to string")
	}

	uidS, ok := uid.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session uid to string")
	}

	emailS, ok := email.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session email to string")
	}

	birthDayI, ok := birthDay.(int)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session birthDay to int")
	}

	birthMonthI, ok := birthMonth.(int)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session birthMonth to int")
	}

	birthYearI, ok := birthYear.(int)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session birthYear to int")
	}

	profilePicS, ok := profilePic.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session profilePic to string")
	}

	user := AuthenticatedUser{
		Username:   usernameS,
		UniqueName:     uniqueNameS,
		Uid:        uidS,
		Email:      emailS,
		BirthDay:   birthDayI,
		BirthMonth: birthMonthI,
		BirthYear:  birthYearI,
		ProfilePic: profilePicS,
	}

	return user, err.Error{}
}

func CreateUserSession(w http.ResponseWriter, r *http.Request, store *redisstore.RedisStore, user AuthenticatedUser) err.Error {
	vErr := user.validateAuthenticatedUser()
	if vErr.E != nil {
		return vErr
	}

	userSession := os.Getenv("USER_SESSION")
	session, sErr := store.Get(r, userSession)
	if sErr != nil {
		return err.NewFromErr(sErr)
	}

	session.Values["username"] = user.Username
	session.Values["uniqueName"] = user.UniqueName
	session.Values["uid"] = user.Uid
	session.Values["email"] = user.Email
	session.Values["birthDay"] = user.BirthDay
	session.Values["birthMonth"] = user.BirthMonth
	session.Values["birthYear"] = user.BirthYear
	session.Values["profilePic"] = user.ProfilePic

	wErr := session.Save(r, w)
	if wErr != nil {
		return err.NewFromErr(wErr)
	}

	return err.Error{}
}

func DestroyUserSession(w http.ResponseWriter, r *http.Request, store *redisstore.RedisStore) err.Error {
	sessionName := os.Getenv("USER_SESSION")
	session, sErr := store.Get(r, sessionName)
	if sErr != nil {
		return err.NewFromErr(sErr)
	}

	session.Options.MaxAge = -1
	dErr := session.Save(r, w)
	if dErr != nil {
		return err.NewFromErr(dErr)
	}

	return err.Error{}
}
