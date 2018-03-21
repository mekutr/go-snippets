package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    
    //testScanf()
    //testBufioOs()
    //testBufioOsNumericValues()
}

// Scan methods are designed to get input from console, files, variables, and
// other input sources and automatically separate values from each other by
// looking for space characters
func testScanf() {
    var s string

    // Scan function is designed for analyzing and parsing the string
    // It automatically breaks up the string whenever it finds space characters

    fmt.Scanln(&s) // --> & character passes the reference of s
    fmt.Println(s)
}

// If you want simply to collect user input in a console application, you should
// instead use a couple of packages bufio and os
func testBufioOs() {

    // a reader object can collect information from a variety of inputs
    reader := bufio.NewReader(os.Stdin) // --> this means that this reader object is looking for information from standard input
    fmt.Print("Enter text: ")
    str, _ := reader.ReadString('\n') // --> the character defines when the ReadString operation needs to return
                                      // --> in this case, it will return whatever it reads until the new line character
    fmt.Println(str)

}

func testBufioOsNumericValues() {

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a number: ")
    str, _ := reader.ReadString('\n')
    // strconv is used for converting string values to different types
    f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // 64 indicates the bitsize of the float number going to be returned
                                                             // TrimSpace will remove the leading or trailing spaces from the string
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Value of number:", f)
    }
}
