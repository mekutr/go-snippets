package main

/*
-Go is organized by packages and packages have functions
-Your own application has its own package, always named main and it also has the main function
 which is called automatically by the run time as the application starts up
-You can define your own functions and you can then organize them into their own packages
*/

import "fmt"

func main() {
	doSomething()

	sum := addValues(23, 54)
	fmt.Println("Sum:", sum)

	sum = addAllValues(12, 54, 79)
	fmt.Println("New sum:", sum)
}

// As long as the function definitions and the code that calls them are in the same package,
// the case of the function name can be either upper or lower

func doSomething() {
	fmt.Println("Doing something!")
}

// the return type is declared after the closing parenthesis -- in this case it is int
// func addValues(value1 value2 int) int { // if the arguments are the same type you can make it more concise
// you can pass the argument list and declare the type once
func addValues(value1 int, value2 int) int {
	return value1 + value2
}

// You can create functions that accepts arbitrary numbers of values as long as they're all of the same type
// as in the following syntax
func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values) // this returns the type of the values -- it return an int slice
	return sum
}
