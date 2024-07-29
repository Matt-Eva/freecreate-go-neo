package test_handlers

import (
	"context"
	"freecreate/internal/api/middleware"
	"freecreate/internal/err"
	"freecreate/internal/queries"
	"net/http"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

type returnUser struct {
	Username string `json:"username"`
	DisplayName string `json:"displayName"`
}

func HandleMasterUser( ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		handleMasterUser(w, r,ctx, neo, store)
	}
}

func handleMasterUser(w http.ResponseWriter, r *http.Request,ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore){
	aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil{
		// user, uErr := getMasterUserFromDb(ctx, neo)
	}
}

func getMasterUserFromDb(ctx context.Context, neo neo4j.DriverWithContext)(returnUser, err.Error){
	dbName := os.Getenv("NEO_DB")
	query := `
		MATCH (u:User)
		WHERE u.masterUser = true
		RETURN u.username AS Username, u.displayName AS DisplayName
	`
	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(dbName))
	if nErr != nil {
		return returnUser{}, err.NewFromErr(nErr)
	}
	if len(result.Records) < 1 {
		return returnUser{}, err.New("no records returned from query")
	}
	resultMap := result.Records[0].AsMap()
	user := returnUser{}
	cErr := queries.NeoRecordToStruct(resultMap, &user)
	if cErr.E != nil {
return returnUser{}, err.Error{}
	}

	return user, err.Error{}
}