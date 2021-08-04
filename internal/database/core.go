package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "mak"
)

func ConnectDB() {
		// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		
		// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	
		// close database
	defer db.Close()

		// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}