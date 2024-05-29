package handlers

import (
	"encoding/json"
	// "fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request, neo, mongo string) {

	type Message struct {
		Neo   string `json:neo`
		Mongo string `json:mongo`
	}

	message := Message{
		neo, mongo,
	}

	json.NewEncoder(w).Encode(message)

}
