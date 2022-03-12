package data

import (
	"database/sql" //go database driver package
	"log"

	//sqlite3 package
	_ "github.com/mattn/go-sqlite3"
)

//package level variable to hold our database instance
var db *sql.DB

//function to open database connection
func OpenDatabase() error {
	//error is only return value
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db") //driver name and location of DB. sql.Open returns DB instance and error
	//if we hit an error we just return it
	if err != nil {
		return err
	}
	//return error we get when we ping the db connection
	return db.Ping()
}

func CreateListenersTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS "Listeners" (
		"LID"	INTEGER NOT NULL,
		"Name"	TEXT NOT NULL,
		"Port"	INTEGER NOT NULL,
		"ActiveConnectedAgents"	INTEGER NOT NULL,
		PRIMARY KEY("LID" AUTOINCREMENT)
	);`

	statement, err := db.Prepare(createTableSQL) //db.Prepare returns a SQL statement and an error
	if err != nil {
		log.Fatal(err.Error()) //log the error if we get it
	}

	statement.Exec()
	log.Println("Listeners table created")
}

func CreateUsersTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS "Users" (
		"UID"	INTEGER NOT NULL,
		"Username"	TEXT NOT NULL,
		"Password"	TEXT NOT NULL,
		"IsAdmin"	INTEGER NOT NULL,
		PRIMARY KEY("UID" AUTOINCREMENT)
	);`

	statement, err := db.Prepare(createTableSQL) //db.Prepare returns a SQL statement and an error
	if err != nil {
		log.Fatal(err.Error()) //log the error if we get it
	}

	statement.Exec()
	log.Println("Users table created")
}

func CreateCommandLogTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS "CommandLog" (
		"CommandID"	INTEGER NOT NULL,
		"UUID"	INTEGER NOT NULL,
		"Result"	INTEGER NOT NULL,
		PRIMARY KEY("CommandID" AUTOINCREMENT)
	);`

	statement, err := db.Prepare(createTableSQL) //db.Prepare returns a SQL statement and an error
	if err != nil {
		log.Fatal(err.Error()) //log the error if we get it
	}

	statement.Exec()
	log.Println("CommandLog table created")
}

func CreateAgentTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS "Agent" (
		"AID"	INTEGER NOT NULL,
		"UUID"	TEXT NOT NULL,
		"User"	TEXT NOT NULL,
		"IP"	TEXT NOT NULL,
		"IsDeleted"	INTEGER NOT NULL,
		"LastCallback"	INTEGER NOT NULL,
		"CallbackInterval"	INTEGER NOT NULL,
		"Jitter"	INTEGER NOT NULL,
		"Listener"	INTEGER NOT NULL,
		PRIMARY KEY("AID" AUTOINCREMENT)
	);`

	statement, err := db.Prepare(createTableSQL) //db.Prepare returns a SQL statement and an error
	if err != nil {
		log.Fatal(err.Error()) //log the error if we get it
	}

	statement.Exec()
	log.Println("Agents table created")
}

func InsertListener(name string, port string, protocol string) {
	InsertListenerSQL := `INSERT INTO Listeners (name, port, protocol)
	VALUES (?, ?, ?)`

	statement, err := db.Prepare(InsertListenerSQL)
	if err != nil { // if we get an error, log it to the console
		log.Fatalln(err)
	}

	_, err = statement.Exec(name, port, protocol) //execute our statement
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted listener successfully")
}