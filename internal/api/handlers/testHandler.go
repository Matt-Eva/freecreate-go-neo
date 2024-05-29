package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request, neo, mongo, redis string) {
	params := r.URL.Query()
	fmt.Println(params)

	type Message struct {
		Neo   string `json:neo`
		Mongo string `json:mongo`
		Redis string `json:redis`
	}

	message := Message{
		neo, mongo, redis,
	}

	json.NewEncoder(w).Encode(message)

}
