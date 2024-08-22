package auth

import (
	"freecreate/internal/middleware"
	"net/http"

	"github.com/rbcervilla/redisstore/v9"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := middleware.DestroyUserSession(w, r, store)
		if e.E != nil {
			e.Log()
			http.Error(w, e.E.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
