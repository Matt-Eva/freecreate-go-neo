package routes

import (
	// "encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"freecreate/internal/api/handlers"
	"freecreate/internal/api/middleware"
)

func CreateRoutes(neo, mongo, redis string) error {
	router := mux.NewRouter()

	router.HandleFunc("/api", middleware.AddDrivers(handlers.TestHandler, neo, mongo, redis))

	return http.ListenAndServe(":8080", router)
}
