package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateRoutes() error {
	router := mux.NewRouter()

	router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		type Message struct {
			Content string `json:"content"`
		}

		message := Message {
			"hello world!",
		}

		json.NewEncoder(w).Encode(message)
	})

	return http.ListenAndServe(":8080", router)
}