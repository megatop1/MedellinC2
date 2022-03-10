package main

import(
    "fmt"
    "log"
    "net/http"
)

func main() {
    fileServer := http.FileServer(http.Dir("./static")) //The := declares and defines a variable. We are telling go we want to look at the directory named "static". We are creating a path.
    http.Handle("/", fileServer) //handle the root route 
    //http.HandleFunc("/listener", listenerHandler)

    fmt.Printf("Starting server at port 8080\n") //Message to print when we start our go server
    if err := http.ListenAndServe(":8080", nil); err != nil { //ListenAndServe function creates the server. This is the heart of the server. We can have an error, or "nil" and when not nil, log the error
		log.Fatal(err)
	} 
}

func listenerHandler(w http.ResponseWriter, r *http.Request) { //In any API, you have a response and a request. User sends requests, and Server sends response. * is a pointer pointing to the request
	if r.URL.Path != "/listener" { //if request URL 
	http.Error(w, "404 not found", http.StatusNotFound)
	return 
	}
	if r.Method != "GET" {//We don't want people to be able to POST to each page. 
		http.Error(w, "method is not supported", http.StatusNotFound)
	}	
	fmt.Fprintf(w, "Welcome!")
}
