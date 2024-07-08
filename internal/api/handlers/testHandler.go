package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/redis/go-redis/v9"
)

func TestHandler(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext, mongo string, redis *redis.Client, ctx context.Context) {
	params := r.URL.Query()
	fmt.Println(params)

	type Message struct {
		Neo   string `json:"neo"`
		Mongo string `json:"mongo"`
		Redis string `json:"redis"`
	}

	message := Message{
		"neo", "mongo", "redis",
	}

	json.NewEncoder(w).Encode(message)

}
