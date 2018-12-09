package main

import "fmt"

func main() {

	n1, l1 := FullName("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, number of chars: %v", n1, l1)

	fmt.Println()

	n2, l2 := FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v, number of chars: %v", n2, l2)

}

// the second parenthesis indicates the return types -- in this case it returns a string and an int
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

// You can also declare the names of return values in the function signature. This has the effect of
// declaring their names and types for use within the function, and then you only have to set their
// values as opposed to returning the values
// this feature called the naked return because you don't declare what you are returning
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
