package utils

import (
	"log"
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/models"
)


func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

//RespondWithError is
func RespondWithError(w http.ResponseWriter, status int, error models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

//ResponseJSON is
func ResponseJSON(w http.ResponseWriter, data interface{}){
	json.NewEncoder(w).Encode(data)
}



//TokenVerify is
func TokenVerify(next http.HandlerFunc) http.HandlerFunc {
	log.Println("Token verify called")
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		var errorObject models.Error

		authHeader := r.Header.Get("Authorization")

		// fmt.Println(authHeader)
		
		bearerToken := strings.Split(authHeader," ")
		
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token)(interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error")
				}
				return []byte("secret"), nil
			})

			if error != nil {
				errorObject.Message = error.Error()
				RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				log.Println("token is valid")
				next.ServeHTTP(w,r)
			} else {
				errorObject.Message = error.Error()
				RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token"
			RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}

//GenerateJWTToken is
func GenerateJWTToken(user models.User) (string, error) {
	var err error

	//create secret (we can assign any string)
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":user.Email,
		"iss":"course",
	})

	tokenString, err := token.SignedString([]byte(secret))
	logFatal(err)
	return tokenString, nil
}