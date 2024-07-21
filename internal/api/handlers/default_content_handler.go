package handlers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func DefaultContentHandler(w http.ResponseWriter, r *http.Request, mongo *mongo.Client) {

}
