package stringutil

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
