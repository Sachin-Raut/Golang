package userrepository

import (
	"log"
	"database/sql"
	"06-GOLANG-REST-JWT-Mike-M/02-JWT-With-refactor/models"
)

var err error

func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

//UserRepository is
type UserRepository struct {}


//Signup is
func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {
	query := "insert into users(email, password) values($1, $2) returning id;"	
	err = db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)	
	logFatal(err)
	return user
}




//Login is
func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, string) {
	row := db.QueryRow("select * from users where email = $1", user.Email)
	err = row.Scan(&user.ID, &user.Email, &user.Password)

	message := ""

	if err != nil {
		if err == sql.ErrNoRows {
			message = "User doesn't exist"				
		} else {
			logFatal(err)
		}
		return user, message
	}
	return user, message
}