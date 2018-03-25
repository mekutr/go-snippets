package main

import(
    "fmt"
    "net/http"
)

/*
    -Go makes it very easy to create a simple HTTP server application.
    -This is only a very simple HTTP server that listens for requests and responds
    -There is a lot more than this to create a full fledged HTTP server than just
     responding to a single request.
     --You need to add code to handle multiple concurrent requests
     --It needs to deal with HTTP header and parameters and so on
    -This only shows how easy it is to work with the web in Go
*/

// this creates a struct type with no fields
type Hello struct {}

// this creates the ServeHTTP method for the type Hello
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request){
    // Fprint function prints to a writer object
    fmt.Fprint(w, "<h1>Hello from the Go web server!!!</h1>")
}

func main() {
    var h Hello
    /*
        ListenAndServe function accepts the previously declared Hello object
        as an object and calls the predefined ServeHTTP function to serve for
        the request
    */
    err := http.ListenAndServe("localhost:8080", h)
    checkError(err)
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}