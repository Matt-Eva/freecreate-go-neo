package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodePostBody(w http.ResponseWriter, body io.ReadCloser, postData interface{}) {
	err := json.NewDecoder(body).Decode(postData)
	if err != nil{
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}