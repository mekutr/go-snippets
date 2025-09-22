package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Go has a rich set of tools that let you work easily with web applications and services
The HTTP package lets you make requests and send data to remote hosts, and also lets you
create HTTP server applications that listen for and respond to requests.

This go snippet can requests and displays some data from a remote host.
*/

func main() {

	// you can change this url to anything you want to
	url := "https://jsonplaceholder.typicode.com/posts"

	fmt.Printf("Started reading...\n[%v]\n", url)

	resp, err := http.Get(url)

	// errors out if anything goes bad at the time of
	// reading it from the remote source
	if err != nil {
		panic(err)
	}

	fmt.Println("...Finished reading.")

	// this is here for debugging purposes to see
	// the response type
	fmt.Printf("Response type: %T\n", resp)

	// this will open an io connection to read the
	// body of the requested end point
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // close the io connection immediately
	// after creating the io call

	if err != nil {
		panic(err)
	}

	// casts the bytes slice to string
	content := string(bytes)
	fmt.Printf("Response body:\n%v", content)

}
