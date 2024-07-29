package middleware

import (
	"fmt"
	"freecreate/internal/err"
	"net/http"
	"os"

	"github.com/rbcervilla/redisstore/v9"
)

type authenticatedUser struct {
	Username string
	DisplayName string
}

func AuthenticateUser(r *http.Request, store *redisstore.RedisStore) (authenticatedUser, err.Error) {
	sessionName := os.Getenv("USER_SESSION")
	userSession, gErr := store.Get(r, sessionName)
	if gErr != nil {
		return authenticatedUser{}, err.NewFromErr(gErr)
	}

	username, displayName := userSession.Values["username"], userSession.Values["displayName"] 
	if username == nil || displayName == nil {
		msg := fmt.Sprintf("user session attribute(s) nil\n: username: %T\n, displayName: %T", username, displayName, )
		fmt.Println(msg)
		return authenticatedUser{}, err.New(msg)
	}

	usernameS, ok := username.(string)
	if !ok {
		return authenticatedUser{}, err.New("could not convert session username to string")
	}

	displayNameS, ok := displayName.(string)
	if !ok {
		return authenticatedUser{}, err.New("could not convert session displayname to string")
	}

	// uidS, ok := uid.(string)
	// if !ok {
	// 	return authenticatedUser{}, err.New("could not convert session uid to string")
	// }

	user := authenticatedUser {
		Username: usernameS,
		DisplayName: displayNameS,
	}

	fmt.Println("user authenticated")

	return user, err.Error{}
}
