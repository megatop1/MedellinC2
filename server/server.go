package main

import(
    "fmt"
    "log"
    "net/http"
    "net"
    "bufio"
)

func main() {
    listener()
    webserver()
}

func webserver() {
    fileServer := http.FileServer(http.Dir("../Web")) //The := declares and defines a variable. We are telling go we want to look at the directory named "static". We are creating a path.
    http.Handle("/", fileServer) //handle the root route 

    fmt.Printf("Starting server at port 8080\n") //Message to print when we start our go server
    if err := http.ListenAndServe(":8080", nil); err != nil { //ListenAndServe function creates the server. This is the heart of the server. We can have an error, or "nil" and when not nil, log the error
		log.Fatal(err)
	} 
}

func listener() {
	//Listen for a connection
    //Accept a connection
    //Handle the connection in a seperate threat to allow multiple clients to connect to this server
    dstream, err := net.Listen("tcp", ":8000") //create a data astream

    if err != nil { //error handling
        fmt.Println(err)
        return 
    }
    defer dstream.Close()//defer data stream from closing so we have opportunity to read stream before connection closes

    for { //infinite loop
        con, err := dstream.Accept() //Once it listens for a connection and its found, it will accept the connection and initiate the connection stream. Handle connection and read data
        if err != nil {
            fmt.Println(err)
            return
        }

        go handle(con) //handle the connection
    }
}

//function to handle the connection 
//Read data out of a connection stream into a buffer, throw it into a variable, run our checks and handle the connections
func handle(con net.Conn) { //tell it that the parameters is conn from the net package. Handles the connection in go routine
    for { //infinite loop
        data, err := bufio.NewReader(con).ReadString('\n') //buffer io package, allows us to start reading data. This is an I/O Reader. Many other functions than ReadString. Only read when we get a \n
        if err != nil {
            fmt.Println(err)
            return
        }
        //print the data we have
        fmt.Println(data)
    }
    //close the connection
    con.Close()

}