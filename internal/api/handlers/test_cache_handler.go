package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func TestCachePostHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client, ctx context.Context) {
	type PostData struct {
		Name string `json:"name"`
	}

	var postData PostData
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	fmt.Println(postData)

	json.NewEncoder(w).Encode(&postData)
}

func TestCacheGetHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client, ctx context.Context) {
	// body := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
}
