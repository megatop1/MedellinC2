package main

import(
    "fmt"
    "log"
    "net/http"
)

func main() {
    fileServer := http.FileServer(http.Dir("./static")) //The := declares and defines a variable. We are telling go we want to look at the directory named "static". We are creating a path.
    http.Handle("/", fileServer) //handle the root route 
    http.HandleFunc("form", formHandler) //function in HTTP package to handle another route
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8080\n") //Message to print when we start our go server
    if err := http.ListenerAndServe(":8080", nil) //
}