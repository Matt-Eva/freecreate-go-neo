package routes

import (
	// "encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"freecreate/internal/api/handlers"
)

func CreateRoutes(neo, mongo string) error {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.TestHandler(neo, mongo))

	return http.ListenAndServe(":8080", router)
}

