package driver

import (
	
	"database/sql"
	_ "github.com/lib/pq" //pq is
	"log"
	"fmt"
)

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

var db *sql.DB
var err error

func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=%s",
	host, port, user, password, dbname, sslmode)
	
	db, err = sql.Open("postgres", connStr)

	logFatal(err)

	err = db.Ping()

	logFatal(err)
	log.Println("Hi")
	return db
}