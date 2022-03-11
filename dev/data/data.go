package data

import (
	"database/sql" //go database driver package

	_ "github.com/mattn/go-sqlite3" //sqlite3 package
)

//package level variable to hold our database instance
var db *sql.DB

//function to open database connection
func OpenDatabase() error {
	//error is only return value
	var err error

	db, err = sql.Open("sqlite3,", "./sqlite-database.db") //driver name and location of DB. sql.Open returns DB instance and error
	//if we hit an error we just return it
	if err != nil {
		return err
	}
	//return error we get when we ping the db connection
	return db.Ping()
}
