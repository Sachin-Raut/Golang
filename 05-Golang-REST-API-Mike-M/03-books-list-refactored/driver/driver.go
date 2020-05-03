package driver

import (
	"database/sql"
	"log"
	"fmt"
)

var db *sql.DB
var err error


func logFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
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


//ConnectDB is
func ConnectDB() *sql.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=%s",
	host, port, user, password, dbname, sslmode)
	
	db, err = sql.Open("postgres", connStr)

	logFatal(err)

	db.Ping()
	return db
}