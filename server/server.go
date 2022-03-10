package main

import(
    "fmt"
    "log"
    "net/http"
)

func main() {
    fileServer := http.FileServer(http.Dir("../Web")) //The := declares and defines a variable. We are telling go we want to look at the directory named "static". We are creating a path.
    http.Handle("/", fileServer) //handle the root route 

    fmt.Printf("Starting server at port 8080\n") //Message to print when we start our go server
    if err := http.ListenAndServe(":8080", nil); err != nil { //ListenAndServe function creates the server. This is the heart of the server. We can have an error, or "nil" and when not nil, log the error
		log.Fatal(err)
	} 
}

