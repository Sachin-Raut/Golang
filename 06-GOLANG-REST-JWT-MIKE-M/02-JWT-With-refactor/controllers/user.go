package controllers

import (
	"log"
	"fmt"
	"database/sql"
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/models"
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/utils"
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/repository/user"
)


func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

//Controller is
type Controller struct {}


//Signup is
func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){

		var user models.User
		var error models.Error
	
		//this will be sent to postman (i.e client)
		// w.Write([]byte("Successfully called signup"))
	
		json.NewDecoder(r.Body).Decode(&user)
		if user.Email == "" {
			error.Message = "Email is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
	
		if user.Password == "" {
			error.Message = " Password is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
	
		//let print user values to terminal
		fmt.Println(user)
	
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	  
		logFatal(err)
		log.Println("pass text - ", user.Password)
		log.Println("hash text - ", hash) //hash is byteArray
	
		//convert byteArray to string
		user.Password = string(hash)
		userRepo := userrepository.UserRepository{}
		user = userRepo.Signup(db, user)
	
		log.Println("after hashing - ", user.Password)
	
		//let's return the user object to client, but don't reveal the hashed password
		user.Password = ""
		w.Header().Set("Content-Type","application/json")
		utils.ResponseJSON(w, user)
	}
}

//Login is
func (c Controller) Login(db *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){
		//this will be sent to postman (i.e client)
		// w.Write([]byte("Successfully called login"))
	
		var user models.User
		var jwt models.JWT
		var error models.Error
		json.NewDecoder(r.Body).Decode(&user)
	
	
		if user.Email == "" {
			error.Message = "Email is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
	
		if user.Password == "" {
			error.Message = "Password is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
	
		//check if the user exists
		password := user.Password

		userRepo := userrepository.UserRepository{}
		user, message := userRepo.Login(db, user)
		
		if message != "" {
			error.Message = "User doesn't exist"
			utils.RespondWithError(w, http.StatusBadRequest, error)			
			return
		}

		log.Println("User returned successfully")

		hashedPassword := user.Password
	
		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	
		if err != nil {
			error.Message = "Invalid Password"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return
		}

		token, err := utils.GenerateJWTToken(user)
		logFatal(err)
		w.WriteHeader(http.StatusOK)
		jwt.Token = token
		utils.ResponseJSON(w, jwt)
	
	}

}


