package main

import (
    "fmt"
)

/*
    -The struct type in Go is a data structure. Structs are similar in purpose and function to C's
     struct and Java's classes. They encapsulate both data and methods.
    -In Go, you don't have concepts like super and sub-structs
    -Each structure is independent. with its own fields for data management and optionally its own methods
*/

// you define a struct as its own custom type. The structs name typically has an initial uppercase character
type Dog struct {
    Breed string
    Weight int
}

func main() {
    poodle := Dog{"Poodle", 34}
    fmt.Println(poodle)
    // you can dump the complete contents of a struct,
    // including its field names and values using the
    // Printf function an a verb spelled %+v
    fmt.Printf("%+v\n", poodle)
    // you can access the structs individual fields with the dot(.) operator
    fmt.Printf("Breed: %v\nWeight: %v", poodle.Breed, poodle.Weight)
}
