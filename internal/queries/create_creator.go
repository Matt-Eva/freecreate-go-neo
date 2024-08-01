package queries

import (
	"context"
	"fmt"
	"freecreate/internal/api/handlers"
	"freecreate/internal/api/middleware"
	"freecreate/internal/err"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateCreator(ctx context.Context, neo neo4j.DriverWithContext, user middleware.AuthenticatedUser, creator handlers.PostedCreator){

}

func checkUniqueCreator(){}

func buildCheckUniqueCreatorQuery()(string, err.Error){

	return "", err.Error{}
}

func buildCheckUniqueCreatorParams(){}

func buildCreateCreatorQuery()(string, err.Error){
	creatorLabel, cErr := GetNodeLabel("Creator")
	if cErr.E != nil {
		return "", cErr
	}

	userLabel, uErr := GetNodeLabel("User")
	if uErr.E != nil {
		return "", uErr
	}

	isCreatorLabel, rErr := GetRelationshipLabel("IS_CREATOR")
	if rErr.E != nil {
		return "", rErr
	}

query := fmt.Sprintf("CREATE (c:%s $creatorParams) <-[r:%s]-(u:%s {uid: $userId})", creatorLabel, isCreatorLabel, userLabel)
return query, err.Error{}
}

func buildCreateCreatorParams(){

}