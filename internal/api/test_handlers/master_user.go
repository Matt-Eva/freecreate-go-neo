package test_handlers

import (
	"context"
	"encoding/json"
	"freecreate/internal/api/middleware"
	"freecreate/internal/err"
	"freecreate/internal/utils"
	"net/http"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/rbcervilla/redisstore/v9"
)

func HandleMasterUser(ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleMasterUser(w, r, ctx, neo, store)
	}
}

func handleMasterUser(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext, store *redisstore.RedisStore) {
	user, aErr := middleware.AuthenticateUser(r, store)
	if aErr.E != nil {
		result, uErr := getMasterUserFromDb(ctx, neo)
		if uErr.E != nil {
			uErr.Log()
			http.Error(w, uErr.E.Error(), http.StatusInternalServerError)
			return
		}

		user := middleware.AuthenticatedUser{}
		utils.MapToStruct(result, &user)

		sErr := middleware.CreateUserSession(w, r, store, user)
		if sErr.E != nil {
			sErr.Log()
			http.Error(w, sErr.E.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func getMasterUserFromDb(ctx context.Context, neo neo4j.DriverWithContext) (map[string]any, err.Error) {
	dbName := os.Getenv("NEO_DB")
	query := `
		MATCH (u:User)
		WHERE u.masterUser = true
		RETURN u.username AS Username, u.userId AS UserId, u.uid AS Uid
	`
	result, nErr := neo4j.ExecuteQuery(ctx, neo, query, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(dbName))
	if nErr != nil {
		return map[string]any{}, err.NewFromErr(nErr)
	}
	if len(result.Records) < 1 {
		return map[string]any{}, err.New("no records returned from query")
	}
	resultMap := result.Records[0].AsMap()

	return resultMap, err.Error{}
}
