package middleware

import (
	"net/http"
)

func Driver(f http.HandlerFunc, driver string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
	}
}
