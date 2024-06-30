package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func TestCachePostHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client) {
	type PostData struct {
		Name string `json:"name"`
	}

	var body PostData
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(body.Name)
}

func TestCacheGetHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client) {
	// body := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
}
