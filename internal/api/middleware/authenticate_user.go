package middleware

import (
	"freecreate/internal/err"
	"net/http"

	"github.com/rbcervilla/redisstore/v9"
)

func AuthenticateUser(r *http.Request, store *redisstore.RedisStore) err.Error {
	_, gErr := store.Get(r, "uid")
	if gErr != nil {
		return err.NewFromErr(gErr)
	}

	return err.Error{}
}
