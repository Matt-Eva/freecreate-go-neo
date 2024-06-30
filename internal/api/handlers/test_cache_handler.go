package handlers

import (
	"encoding/json"
	"fmt"
	"freecreate/internal/utils"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func TestCachePostHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client) {
	type PostData struct {
		Name string `json:"name"`
	}

	var postData PostData
	utils.DecodePostBody(w, r.Body, &postData)
	

	json.NewEncoder(w).Encode(&postData)
}

func TestCacheGetHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client) {
	// body := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
}
