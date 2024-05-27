package handlers

import (
	"encoding/json"
	// "fmt"
	"net/http"
)

func TestHandler(neo, mongo string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Message struct {
			Neo   string `json:neo`
			Mongo string `json:mongo`
		}

		message := Message{
			neo, mongo,
		}

		json.NewEncoder(w).Encode(message)
	}
}
