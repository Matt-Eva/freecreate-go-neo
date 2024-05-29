package middleware

import (
	"net/http"
)

func AddDrivers(handler func(w http.ResponseWriter, r *http.Request, neo, mongo string), neo, mongo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, neo, mongo)
	}
}
