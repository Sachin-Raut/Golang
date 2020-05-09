package main

import (
	
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"strings"

	// "05-Golang-REST-API-Mike-M/03-books-list-refactored/driver"
	// "05-Golang-REST-API-Mike-M/03-books-list-refactored/controllers"
)

func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

//User is
type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

//JWT is
type JWT struct {
	Token string `json:"token"`
}

//Error is
type Error struct {
	Message string `json:"message"`
}

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

const (
	//sslmode can be disable or verify-full
	host     = "satao.db.elephantsql.com"
	port     = 5432
	user     = "bgligpzc"
	password = "cm_uF7PbzYycl39koGWlaCLAEd6mnuF9"
	dbname   = "bgligpzc"
	sslmode  = "verify-full"
	// sslmode  = "disable" 
  )

var db *sql.DB
var err error

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=%s",
	host, port, user, password, dbname, sslmode)
	
	db, err = sql.Open("postgres", connStr)

	logFatal(err)

	db.Ping()
	//create router instance
	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")

	router.HandleFunc("/login", login).Methods("POST")

	router.HandleFunc("/protected", TokenVerify(protectedEndpoint)).Methods("GET")

	//start server
	fmt.Println("Started server")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//signup is
func signup(w http.ResponseWriter, r *http.Request){

	var user User
	var error Error

	//this will be sent to postman (i.e client)
	// w.Write([]byte("Successfully called signup"))

	json.NewDecoder(r.Body).Decode(&user)
	if user.Email == "" {
		error.Message = "Email is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = " Password is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	//let print user values to terminal
	fmt.Println(user)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
  
	if err != nil {
		logFatal(err)
	}
	log.Println("pass text - ", user.Password)
	log.Println("hash text - ", hash) //hash is byteArray

	//convert byteArray to string
	user.Password = string(hash)

	log.Println("after hashing - ", user.Password)

	query := "insert into users(email, password) values($1, $2) returning id;"
	
	err = db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		log.Println(err)
		error.Message = "Server error"
		respondWithError(w, http.StatusInternalServerError, error)
		return
	}

	//let's return the user object to client, but don't reveal the hashed password
	user.Password = ""
	w.Header().Set("Content-Type","application/json")
	responseJSON(w, user)
}

func responseJSON(w http.ResponseWriter, data interface{}){
	json.NewEncoder(w).Encode(data)
}

//login is
func login(w http.ResponseWriter, r *http.Request){
	//this will be sent to postman (i.e client)
	// w.Write([]byte("Successfully called login"))

	var user User
	var jwt JWT
	var error Error
	json.NewDecoder(r.Body).Decode(&user)


	if user.Email == "" {
		error.Message = "Email is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	//check if the user exists
	password := user.Password

	row := db.QueryRow("select * from users where email = $1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "User doesn't exist"
			respondWithError(w, http.StatusBadRequest, error)			
		} else {
			logFatal(err)
		}
		return
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		error.Message = "Invalid Password"
		respondWithError(w, http.StatusUnauthorized, error)
		return
	}

	token, err := GenerateJWTToken(user)
	logFatal(err)
	w.WriteHeader(http.StatusOK)
	jwt.Token = token
	responseJSON(w, jwt)

}

//protectedEndpoint is
func protectedEndpoint(w http.ResponseWriter, r *http.Request){
	log.Println("pppppppppp")
}

//TokenVerify is
func TokenVerify(next http.HandlerFunc) http.HandlerFunc {
	log.Println("Token verify called")
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		var errorObject Error

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
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				log.Println("token is valid")
				next.ServeHTTP(w,r)
			} else {
				errorObject.Message = error.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token"
			respondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}

//GenerateJWTToken is
func GenerateJWTToken(user User) (string, error) {
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