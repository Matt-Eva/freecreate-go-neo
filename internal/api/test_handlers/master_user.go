package test_handlers

import (
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func HandleMasterUser(w http.ResponseWriter, r *http.Request, neo neo4j.DriverWithContext){
	// query := `
	// 	MATCH (u:User)
	// 	WHERE u.masterUser = true
	// 	RETURN u.username AS username, u.uid AS Uid, u.displayName AS displayName
	// `
}