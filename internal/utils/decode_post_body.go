package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodePostBody(w http.ResponseWriter, body io.ReadCloser, postData interface{}) error {
	err := json.NewDecoder(body).Decode(postData)
	if err != nil{
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return err
	}

	return nil
}