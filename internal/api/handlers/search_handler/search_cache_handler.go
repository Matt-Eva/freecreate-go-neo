package search_handler

import (
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func SearchCacheHandler(w http.ResponseWriter, r *http.Request, redis *redis.Client, ctx context.Context) {

}
