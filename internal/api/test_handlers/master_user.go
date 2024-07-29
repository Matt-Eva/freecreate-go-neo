package test_handlers

import (
	"context"
	"encoding/json"
	"freecreate/internal/api/middleware"
	"freecreate/internal/err"
	"freecreate/internal/queries"
	"net/http"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

type returnUser struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

func HandleMasterUser(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleMasterUser(w, r, ctx, neo, store)
	}
}

func handleMasterUser(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	authenticatedUser, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		user, uErr := getMasterUserFromDb(ctx, neo)
		if uErr.E != nil {
			uErr.Log()
			http.Error(w, uErr.E.Error(), http.StatusInternalServerError)
			return
		}

		sErr := createUserSession(w, r, store, user)
		if sErr.E != nil {
			sErr.Log()
			http.Error(w, sErr.E.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}

	user := returnUser{
		Username: authenticatedUser.Username,
		DisplayName: authenticatedUser.DisplayName,
	}
	
	json.NewEncoder(w).Encode(user)
}

func getMasterUserFromDb(ctx context.Context, neo neo4j.DriverWithContext) (returnUser, err.Error) {
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

func createUserSession(w http.ResponseWriter, r *http.Request, store *redisstore.RedisStore, user returnUser) err.Error{
	userSession := os.Getenv("USER_SESSION")
	session, sErr := store.Get(r, userSession)
	if sErr != nil {
		return err.NewFromErr(sErr)
	}

	session.Values["username"] = user.Username
	session.Values["displayName"] = user.DisplayName
	
	wErr := session.Save(r, w)
	if wErr != nil {
		return err.NewFromErr(wErr)
	}

	return err.Error{}
}
