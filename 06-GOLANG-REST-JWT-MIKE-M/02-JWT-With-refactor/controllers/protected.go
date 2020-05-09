package controllers

import (
	"log"
	"net/http"
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/utils"
)

//ProtectedEndpoint is
func (c Controller) ProtectedEndpoint() http.HandlerFunc{

		return func (w http.ResponseWriter, r *http.Request) {
		log.Println("pppppppppp")
		utils.ResponseJSON(w, "Yes")
	}
}
