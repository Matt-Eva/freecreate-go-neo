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
	UserId     string `json:"userId"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	BirthDay   int    `json:"birthday"`
	BirthYear  int    `json:"birthYear"`
	BirthMonth int    `json:"birthMonth"`
	ProfilePic string `json:"profilePic"`
}

func AuthenticateUser(r *http.Request, store *redisstore.RedisStore) (AuthenticatedUser, err.Error) {
	sessionName := os.Getenv("USER_SESSION")
	userSession, gErr := store.Get(r, sessionName)
	if gErr != nil {
		return AuthenticatedUser{}, err.NewFromErr(gErr)
	}

	uid := userSession.Values["uid"]
	userId := userSession.Values["userId"]
	username := userSession.Values["username"]
	email := userSession.Values["email"]
	birthDay := userSession.Values["birthDay"]
	birthMonth := userSession.Values["birthMonth"]
	birthYear := userSession.Values["birthYear"]
	profilePic := userSession.Values["profilePic"]

	if username == nil || userId == nil || uid == nil || email == nil || birthDay == nil || birthMonth == nil || birthYear == nil || birthYear == profilePic {
		msg := fmt.Sprintf(
			"user session attribute(s) nil\n: username: %T\n userId: %T\n uid: %T\n email: %T\n birthDay: %T\n birthMonth: %T\n birthYear: %T\n profilePic: %T",
			username, userId, uid, email, birthDay, birthMonth, birthYear, profilePic)
		return AuthenticatedUser{}, err.New(msg)
	}

	usernameS, ok := username.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session username to string")
	}

	userIdS, ok := userId.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session userId to string")
	}

	uidS, ok := uid.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session uid to string")
	}

	user := AuthenticatedUser{
		Username: usernameS,
		UserId:   userIdS,
		Uid:      uidS,
	}

	fmt.Println("user authenticated")

	return user, err.Error{}
}

func CreateUserSession(w http.ResponseWriter, r *http.Request, store *redisstore.RedisStore, user AuthenticatedUser) err.Error {
	userSession := os.Getenv("USER_SESSION")
	session, sErr := store.Get(r, userSession)
	if sErr != nil {
		return err.NewFromErr(sErr)
	}

	session.Values["username"] = user.Username
	session.Values["userId"] = user.UserId
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
