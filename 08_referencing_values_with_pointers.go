package main

import "fmt"

func main() {

	// this is how you define a pointer in Go
	var p *int

	if p != nil {
		// this is how you deref a pointer -- by putting * in front of a variable
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var v int = 42
	p = &v // this takes the address of v and assigns to p

	if p != nil {
		// this is how you deref a pointer -- by putting * in front of a variable
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 42.13
	pointer1 := &value1
	fmt.Println("Value1:", *pointer1)

	// this example proves that a pointer points to the location of a variable
	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1:", *pointer1)
	fmt.Println("Value 1:", value1)
}
