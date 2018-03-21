package main

/*
Unlike C# or Java, the string type is not a complex type. It does not have bunch of methods that you can use
to manipulate strings. Instead, if you want to manipulate or inspect a string, you can use the functions
that are available in the strings package
*/

import (
    "fmt"
    "strings"
)
func main() {

    str1 := "An implictly typed string"
    fmt.Printf("str1: %v:%T\n", str1, str1)

    fmt.Println(strings.ToUpper(str1))

    fmt.Println(strings.Title(str1)) // changes the first character of each word in a string to uppercase

    lValue := "hello"
    uValue := "HELLO"
    fmt.Println("Equal?", (lValue == uValue)) // == comparison operator compares the contents of the strings
    fmt.Println("Equal Non-Case-Sensitive?", strings.EqualFold(lValue, uValue)) // EqualFold folds all the values to all uppercase

    fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
    fmt.Println("Contains exp?", strings.Contains(str1, "imp"))

}
