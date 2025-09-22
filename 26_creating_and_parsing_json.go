package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
Go has a JSON package that lets you easily create and read JSON formatted text
*/

// this struct should reflect the values that are going to be read from a json file
// please check the readFromJSON function to see the usage of it
type jsonPack struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func main() {
	// you can change this url to anything you want to
	url := "https://jsonplaceholder.typicode.com/posts"
	content := contentFromServer(url)
	rfjson := readFromJSON(content)

	for i := 0; i < len(rfjson); i++ {
		fmt.Println("Title:", rfjson[i].Title) // this will only read the Title
	}
}

// returns the contents of a response after
// requesting the given url as a string
func contentFromServer(url string) string {

	resp, err := http.Get(url)
	checkError(err)

	// this will open an io connection to read the
	// body of the requested end point
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // close the io connection immediately
	// after creating the io call
	checkError(err)

	// casts the bytes slice to string
	return string(bytes)
}

func readFromJSON(content string) []jsonPack {

	decoder := json.NewDecoder(strings.NewReader(content))

	// this returns the next token from a given json string,
	_, err := decoder.Token()
	checkError(err)
	// calling decoder.Token() will remove the first token it encounters.
	// it will end up removing the outer array if you call the Token() method on it
	// the following json packet will end up being in the form of {...}, {...}
	/*
	   [
	    {
	      "userId":"10",
	      "title":"manager"
	    },
	    {
	       "userId":"10",
	       "title":"manager"
	    }
	   ]
	*/

	var jp jsonPack
	var jps []jsonPack
	for decoder.More() {
		/*
		-the decode method is going to parse the json content and grab the fields in the struct.
		-it is going to get only those values that are available in the struct while parsing and
		 putting into struct.
		-the types definition of each field coming from the json needs to match with the types
		 coming from the struct, otherwise it errors out.
	 	-for example
		 if a given json is like this --> {"userId":"1"}
		 and the struct type is like this --> type jsonPack struct { userId int }
		 this will error out because "1" is a string value not an integer value
		*/
		err := decoder.Decode(&jp)
		checkError(err)

		jps = append(jps, jp)
	}

	return jps
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
