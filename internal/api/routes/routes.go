package routes

import (
	"encoding/json"
	"fmt"
	"freecreate/internal/api/handlers"
	"freecreate/internal/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoutes(neo, mongo, redis string) error {
	router := mux.NewRouter()

	router.HandleFunc("/api", middleware.AddDrivers(handlers.TestHandler, neo, mongo, redis))
	router.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request){
		params := r.URL.Query()
		fmt.Println(params)
		json.NewEncoder(w).Encode(params)
	})

	return http.ListenAndServe(":8080", router)
}
