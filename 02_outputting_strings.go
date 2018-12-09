package main

// This package implements a number of functions you can use to format and print strings
// to a variety of output targets
// Print and PrintLine functions print to the console but there are other functions
// that save formatted strings to variables, write them to files and so on

import "fmt"

func main() {
	str1 := "The quick red box"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	// Print and Prinln functions take their values and concatenate them together
	// into a single string seperating the values with a space character and not with commas
	//--
	// This function returns the number of characters in the resulting string and potentially
	// an error object. -- This is an interesting capability of Go functions
	//                  -- They're allowed to return more than one value simultaneously
	// := --> this is the inferred type assignment operator
	stringLength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String Length:", stringLength)
	}

	// there are circumstances where you're getting values back from a function and you know
	// you're not going to get an error object back
	// But the function is going to return that error object and the syntax rules of Go say that
	// if you have a value that you've declared, you must address it in your code

	stringLength2, _ := fmt.Println(str1, str2, str3) // --> _ means that I know I am getting the variable back but I don't want to address it
	fmt.Println("String Length:", stringLength2)

	// this function lets you create strings that have placeholders known as verbs, that defines how values
	// will be formatted
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)

	// This means give me a floating point number with 2 decimal points
	// Because that verb requires an actual floating value and you don't get implicit conversion in Go,
	// We have to explicitly convert the number here
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))

	// We can also use the Printf function to get the types of each variable
	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n",
		str1, str2, str3, aNumber, isTrue)

	// Each of these functions that outputs to the console has versions for outputting
	// to a string or to other targets
	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T",
		str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
