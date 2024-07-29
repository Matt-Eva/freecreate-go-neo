package middleware

import (
	"fmt"
	"freecreate/internal/err"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v9"
)

type authenticatedUser struct {
	Username string
	DisplayName string
	Uid string
}

func AuthenticateUser(r *http.Request, store *redisstore.RedisStore) (*sessions.Session, err.Error) {
	userSession := os.Getenv("USER_SESSION")
	user, gErr := store.Get(r, userSession)
	if gErr != nil {
		return user, err.NewFromErr(gErr)
	}

	username := user.Values["username"]
	if username == nil {
		fmt.Println("user not authenticated")
		return user, err.New("user not logged in")
	}

	fmt.Println("user authenticated")

	return user, err.Error{}
}
