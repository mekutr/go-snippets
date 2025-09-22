package main

import "fmt"

// Go is a statically typed language
// You have to set the type of each variable during the compilation process
// and you can not change them at runtime
// With Go syntax you declare types either explicitly or implicitly
func main() {

	// Explicit typing
	var firstInteger int // the default value is 0 in this case
	var secondInteger int = 42

	var firstString = "This is Go!!!"

	// Implicit typing
	// the varible type are being inferred based on the initial values that are assigned
	thirdInteger := 42
	secondString := "This is Go!!!"

	// constants in Go with Explicit typing
	const cInteger int = 1
	// constants in Go with Implicit typing
	const cString = "This is Go!!!"

	fmt.Printf(`
		First Integer: %d
		Second Integer: %d
		First String: %s
		Third Integer: %d
		Second String: %s`,
		firstInteger, secondInteger, firstString, thirdInteger, secondString)

	// fixed integer types
	/*
	   uint8, uint16, uint32, uint64
	   int8, int16, int32, int64
	*/
	// integer aliases
	/*
	   byte
	   uint  -- depending on the os it might reflect the uint32 or uint64
	   int   -- depending on the os it might reflect the int32 or int64
	   uintptr
	*/

	// float
	/*
	   float32, float64
	*/

	//complex numbers
	/*
	   complex64, complex128
	*/

	//Data collection types
	/*
	   Arrays, Slices, Maps, Structs, etc.
	*/

	//Language organization
	/*
	   Functions -- Function is a type in Go and that's what it makes it to pass as an argument
	   Interfaces
	   Channels
	*/

	//Data management
	/*
	   Pointers -- these are reference variables that points an address in memory
	*/
}
