package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func main() {

    content := "Hello from Go!"

    file, err := os.Create("./temp/fromString.txt")
    checkError(err)
    // this insures file will be closed no matter what else happens throughout the rest of the main function
    defer file.Close()

    ln, err := io.WriteString(file, content)
    checkError(err)

    fmt.Printf("All done with file of %v characters", ln)

    // there is also a function name WriteFile, that writes a byte array instead of
    // a simple string
    bytes := []byte(content)
    ioutil.WriteFile("./temp/fromBytes.txt", bytes, 0644)
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}