package data

import (
	"database/sql" //go database driver package
	"log"
	"strconv"

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
		"Protocol" TEXT NOT NULL,
		"IP" TEXT NOT NULL,
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
		"RemoteIP"	TEXT NOT NULL,
		"Hostname" TEXT NOT NULL,
		"IsAlive" INTEGER NOT NULL DEFAULT 1 CHECK(IsAlive IN (0,1)),
		"Command" TEXT,
		"DefaultDelay" INTEGER NOT NULL,
		"LastTimeCommandExecuted" TEXT,
		"LastTimeCommandWasExecuted" TEXT,
		"TimeToSendNextCommand" TEXT,
		"CurrentTime" TEXT,
 		PRIMARY KEY("AID" AUTOINCREMENT)
	);`

	statement, err := db.Prepare(createTableSQL) //db.Prepare returns a SQL statement and an error
	if err != nil {
		log.Fatal(err.Error()) //log the error if we get it
	}

	statement.Exec()
	log.Println("Agents table created")
}

func CreateLaunchersTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS "Launchers" (
		"LauncherID" INTEGER NOT NULL,
		"RemoteIP"	TEXT NOT NULL,
		"Listener"	TEXT NOT NULL,
		"ListenerIP" TEXT NOT NULL,
		"RemotePort" TEXT NOT NULL,
		"Jitter"	TEXT NOT NULL,
		"PayloadType" TEXT NOT NULL,
		PRIMARY KEY("LauncherID" AUTOINCREMENT)
	);`

	statement, err := db.Prepare(createTableSQL) //db.Prepare returns a SQL statement and an error
	if err != nil {
		log.Fatal(err.Error()) //log the error if we get it
	}

	statement.Exec()
	log.Println("CommandLog table created")
}

func InsertListener(name string, port string, IP string, protocol string) {
	//randomly generate a LID (Listeners Unique ID)

	InsertListenerSQL := `INSERT INTO Listeners (Name, Port, Protocol, IP, ActiveConnectedAgents)
	VALUES (?, ?, ?, ?, 0)`

	statement, err := db.Prepare(InsertListenerSQL)
	if err != nil { // if we get an error, log it to the console
		log.Fatalln(err)
	}

	_, err = statement.Exec(name, port, protocol, IP) //execute our statement
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully Wrote Listener to Database, Proceeding to Generate Listener...")
}

func DisplayAllListeners() {
	row, err := db.Query("SELECT * FROM Listeners")
	if err != nil {
		log.Fatalln(err) //log error if it occurs to the console
	}
	//close the row once we reach end of the function
	defer row.Close()

	//run through all the rows and print them out to the terminal
	for row.Next() {
		var LID int
		var Name string
		var Port int
		var Protocol string
		var IP string
		var ActiveConnectedAgents int
		var DefaultDelay int

		err = row.Scan(&LID, &Name, &Port, &Protocol, &IP, &ActiveConnectedAgents)
		if err != nil { //if there is an issue scanning the row print this error to the console
			log.Fatalln(err)
		}
		log.Println("Listener Name:", Name, "|", IP, "| Port:", Port, "| Connected Agents:", ActiveConnectedAgents, "| DefaultDelay: ", DefaultDelay)
	}
}

func GetIP() string {
	var IP string
	err := db.QueryRow("SELECT IP FROM Listeners ORDER BY LID LIMIT 1").Scan(&IP) //Grab IP from the DB
	if err != nil {
		log.Fatalln(err) //log error if it occurs to the console
	}
	//return the IP
	return IP
}

func GetPort() string {
	var port int
	err := db.QueryRow("SELECT Port FROM Listeners ORDER BY LID DESC LIMIT 1").Scan(&port) //Grab Port from the DB
	if err != nil {
		log.Fatalln(err) //log error if it occurs to the console
	}
	//convert port integer to string value
	strPort := strconv.Itoa(port)
	//return the Port
	return strPort
}

func GetListenerPorts() string {
	var portList string
	err := db.QueryRow("SELECT group_concat(Port, ',') FROM Listeners").Scan(&portList) //Grab Port from the DB
	if err != nil {
		log.Fatalln(err) //log error if it occurs to the console
	}

	//return the Port
	return portList
}

func InsertLauncher(RemoteIP string, Listener string, ListenerIP string, RemotePort string, Jitter string, PayloadType string) {
	InsertLauncherSQL := `INSERT INTO Launchers (RemoteIP, Listener, ListenerIP, RemotePort, Jitter, PayloadType)
	VALUES (?, ?, ?, ?, ?, ?)`

	statement, err := db.Prepare(InsertLauncherSQL)
	if err != nil { // if we get an error, log it to the console
		log.Fatalln(err)
	}

	_, err = statement.Exec(RemoteIP, Listener, ListenerIP, RemotePort, Jitter, PayloadType) //execute our statement
	if err != nil {
		log.Fatalln(err)
	}
}

func InsertAgent(UUID string, RemoteIP string, Hostname string) {
	InsertAgentSQL := `INSERT INTO Agent (UUID, RemoteIP, Hostname, DefaultDelay, CurrentTime, TimeToSendNextCommand)
	VALUES (?, ?, ?, 10, datetime(), datetime('now', '+.1 Minute') )`
	//default DefaultDelay value will be 10 seconds
	statement, err := db.Prepare(InsertAgentSQL)
	if err != nil { // if we get an error, log it to the console
		log.Fatalln(err)
	}

	_, err = statement.Exec(UUID, RemoteIP, Hostname) //execute our statement
	if err != nil {
		log.Fatalln(err)
	}

	/* TimeToSendNextCommand (SendNextCommand = CurrentTime + (DefaultDelay * Jitter)) */
	//db.Query("SELECT UUID FROM Agent WHERE IsAlive=1")
}

func CheckDuplicateAgentUUID() bool {
	var flag bool

	err := db.QueryRow("SELECT UUID, COUNT(*) c FROM Agent GROUP BY UUID HAVING c > 1;")
	if err != nil {
		print("UUID Does Not Repeat\n" + "Generating Agent...\n")
		//log.Fatalln(err) //log error if it occurs to the console
		flag = true
	} else {
		print("UUID Repeats")
		flag = false
	}
	return flag
}

func GetAliveAgents() {
	var uuid string
	print("-----------Alive Agents-----------\n")
	rows, err := db.Query("SELECT UUID FROM Agent WHERE IsAlive=1")
	if err != nil {
		log.Fatalln(err) //log error if it occurs to the console
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&uuid)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(uuid)
		println(uuid)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func GetAgentInformation() {
	row, err := db.Query("SELECT * FROM Agents WHERE UUID =%s")
	if err != nil {
		log.Fatalln(err) //log error if it occurs to the console
	}
	//close the row once we reach end of the function
	defer row.Close()

	//run through all the rows and print them out to the terminal
	for row.Next() {
		var LID int
		var Name string
		var Port int
		var Protocol string
		var IP string
		var ActiveConnectedAgents int

		err = row.Scan(&LID, &Name, &Port, &Protocol, &IP, &ActiveConnectedAgents)
		if err != nil { //if there is an issue scanning the row print this error to the console
			log.Fatalln(err)
		}
		log.Println("Listener Name:", Name, "|", IP, "| Port:", Port, "| Connected Agents:", ActiveConnectedAgents)
	}

}

func GetAgentUUID(uuid string) {
	row, err := db.Query("SELECT * FROM Agent WHERE UUID = \" " + uuid + " \" ")
	if err != nil {
		log.Fatalln(err) //log error if it occurs to the console
	} else {
		print("Agent is in the database\n")
	}
	//close the row once we reach end of the function
	defer row.Close()

	for row.Next() {
		var UUID string

		err = row.Scan(&UUID)
		log.Println("UUID Exists!\n")
	}
}

/* Insert's the User's Command to the Command column in the Agent table in the Database */
func InsertCommandToAgentTableInDB(command string, uuid string) {
	statement, err := db.Prepare("UPDATE Agent SET Command =? WHERE UUID=?")
	if err != nil { // if we get an error, log it to the console
		log.Fatalln(err)
	}

	_, err = statement.Exec(command, uuid) //execute our statement
	if err != nil {
		log.Fatalln(err)
	}
}

func awaitCommands() {
	/* for (DefaultDelayValue) { { */
	/* Loop through every row based off of UUID in the DB */
	/* Checks DefaultDelay value in Agent*/
	/* Check Command section in Agent table for that UUID */
	/* Send the command to the server */

	/* Step 1: Loop through every agent in the Agent table*/
	//row, err := db.Query("SELECT * FROM Agent WHERE CurrentTime = TimeToSendNextCommand")
}

func GetUserCommandFromDB() {
	var uuid string
	print("-----------Alive Agents-----------\n")
	rows, err := db.Query("SELECT UUID FROM Agent WHERE IsAlive=1")
	if err != nil {
		log.Fatalln(err) //log error if it occurs to the console
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&uuid)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(uuid)
		println(uuid)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
