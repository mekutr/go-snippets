// First statement in a file is a package declaration
// Your startup file always has a package of main
// The package name does not have to match the name of the directory or the source code file
package main

import "fmt"

// If you are going to be importing more than one package
// you can use the following format
/*
import (
	"fmt"
	"strings"
)
*/

// The main package should have a function named main
// It should be all lower case and does not have any arguments
func main() {
	fmt.Println("Hello from Go!")
}
