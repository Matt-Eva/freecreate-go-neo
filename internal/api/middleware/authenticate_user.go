package middleware

import (
	"fmt"
	"freecreate/internal/err"
	"net/http"
	"os"

	"github.com/rbcervilla/redisstore/v9"
)

type AuthenticatedUser struct {
	Username string
	DisplayName string
	Uid string
}

func AuthenticateUser(r *http.Request, store *redisstore.RedisStore) (AuthenticatedUser, err.Error) {
	sessionName := os.Getenv("USER_SESSION")
	userSession, gErr := store.Get(r, sessionName)
	if gErr != nil {
		return AuthenticatedUser{}, err.NewFromErr(gErr)
	}

	username, displayName, uid := userSession.Values["username"], userSession.Values["displayName"], userSession.Values["uid"] 
	if username == nil || displayName == nil || uid == nil {
		msg := fmt.Sprintf("user session attribute(s) nil\n: username: %T\n displayName: %T\n uid: %T", username, displayName, uid)
		fmt.Println(msg)
		return AuthenticatedUser{}, err.New(msg)
	}

	usernameS, ok := username.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session username to string")
	}

	displayNameS, ok := displayName.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session displayname to string")
	}

	uidS, ok := uid.(string)
	if !ok {
		return AuthenticatedUser{}, err.New("could not convert session uid to string")
	}

	user := AuthenticatedUser {
		Username: usernameS,
		DisplayName: displayNameS,
		Uid: uidS,
	}

	fmt.Println("user authenticated")

	return user, err.Error{}
}

func CreateUserSession(w http.ResponseWriter, r *http.Request, store *redisstore.RedisStore, user AuthenticatedUser) err.Error{
	userSession := os.Getenv("USER_SESSION")
	session, sErr := store.Get(r, userSession)
	if sErr != nil {
		return err.NewFromErr(sErr)
	}

	session.Values["username"] = user.Username
	session.Values["displayName"] = user.DisplayName
	session.Values["uid"] = user.Uid
	
	wErr := session.Save(r, w)
	if wErr != nil {
		return err.NewFromErr(wErr)
	}

	return err.Error{}
}
