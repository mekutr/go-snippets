package main

import "fmt"

func main() {

	var resultx string
	var resulty string

	var x float64 = 42

	// Go does not require parenthesis around the boolean condition
	if x < 0 {
		resultx = "Less than zero"
	} else if x == 0 {
		resultx = "Equal to zero"
	} else {
		resultx = "Greater than zero"
	}

	fmt.Println("Resultx:", resultx)

	// In Go, you can include an initial statement as part of the if declaration
	// Any variables you declare will then be local to the if block and the memory
	// they use will be eligible for garbage collection after the conditional logic
	// is complete
	if y := 0; y < 0 {
		resulty = "Less than zero"
	} else if y == 0 {
		resulty = "Equal to zero"
	} else {
		resulty = "Greater than zero"
	}

	fmt.Println("Resulty:", resulty)

}
