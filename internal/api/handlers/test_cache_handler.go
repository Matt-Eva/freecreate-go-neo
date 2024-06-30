package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"freecreate/internal/utils"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func TestCachePostHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client, ctx context.Context) {
	type PostData struct {
		Name string `json:"name"`
	}

	var postData PostData
	utils.DecodePostBody(w, r.Body, &postData)
	

	json.NewEncoder(w).Encode(&postData)
}

func TestCacheGetHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client, ctx context.Context) {
	// body := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
}
