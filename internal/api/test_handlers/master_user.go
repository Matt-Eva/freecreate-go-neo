package test_handlers

import (
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

func HandleMasterUser( neo neo4j.DriverWithContext, store *redisstore.RedisStore)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		handleMasterUser(w, r, neo, store)
	}
}

func handleMasterUser(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext, store *redisstore.RedisStore){
	
}