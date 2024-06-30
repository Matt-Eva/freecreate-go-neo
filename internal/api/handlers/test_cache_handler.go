package handlers

import (
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
	

	fmt.Println(postData.Name)
}

func TestCacheGetHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client) {
	// body := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
}
