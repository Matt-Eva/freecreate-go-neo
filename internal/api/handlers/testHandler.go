package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func TestHandler(w http.ResponseWriter, r *http.Request, neo, mongo string, redis *redis.Client) {
	params := r.URL.Query()
	fmt.Println(params)

	type Message struct {
		Neo   string `json:neo`
		Mongo string `json:mongo`
		Redis string `json:redis`
	}

	message := Message{
		neo, mongo, "redis",
	}

	json.NewEncoder(w).Encode(message)

}
