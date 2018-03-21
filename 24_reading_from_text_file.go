package main

import (
    "fmt"
    "io/ioutil"
)

func main() {

    fileName := "./temp/hello.txt"

    // ReadFile return a byte array
    content, err := ioutil.ReadFile(fileName)
    checkError(err)

    result := string(content)

    fmt.Println("Read from file:", result)
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}