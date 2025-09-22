package main

import (
	"errors"
	"fmt"
	"os"
)

/*
-Unlike most C-style languages, Go doesn't support classic exception handling syntax
 with the usual try and catch keywords, and since there isn't any type inheritance
 the concept of a super-type such as exception and subtypes for different kinds of exceptions
 doesn't fit in this language either
-Instead, an error in Go is an instance of an interface that defines a single method, named
 error and that method returns a string and that string is the error message
*/

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		// both of these lines create the same output
		// fmt.Println(err) // Go is just taking a little bit of a shortcut here
		fmt.Println(err.Error()) // this is a slightly verbose approach
		// the error message object always has a method named Error
	}

	// you can create your own Error objects, and all you have to do is override the string
	myError := errors.New("My error string")
	fmt.Println(myError)

	// You can handle error conditions quite simply in some cases using what's called
	// comma-OK syntax.
	// Below is an example for the comma-OK syntax

	attendance := map[string]bool{"Ann": true, "Mike": true}

	// In order to find out whether there is an item in the map that's associated with a
	// particular key, you can use this syntax
	// if there is an item for the given key, the ok variable becomes true and the attended
	// variable will have the associated value from the key
	attended, ok := attendance["Mike"]
	if ok {
		fmt.Println("Mike attended?", attended)
	} else {
		fmt.Println("No info for Mike")
	}

}
