package handlers

import (
	"context"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetCreator(w http.ResponseWriter, r *http.Request) {

}

type PostedCreator struct {
	CreatorName string `json:"creatorName"`
	CreatorId string `json:"creatorId"`
}

func CreateCreator(ctx context.Context, neo neo4j.DriverWithContext) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		createCreator(w, r, ctx, neo)
	}
}

func createCreator(w http.ResponseWriter, r *http.Request, ctx context.Context, neo neo4j.DriverWithContext){

}

func UpdateCreator(w http.ResponseWriter, r *http.Request) {

}

func DeleteCreator(w http.ResponseWriter, r *http.Request) {

}
