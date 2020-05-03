package utils

import (
	"net/http"
	"encoding/json"
	"05-Golang-REST-API-Mike-M/03-books-list-refactored/model"
)

//SendError is
func SendError(w http.ResponseWriter, status int, err model.Error){
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

//SendSuccess is
func SendSuccess(w http.ResponseWriter, data interface{}){
	json.NewEncoder(w).Encode(data)
}