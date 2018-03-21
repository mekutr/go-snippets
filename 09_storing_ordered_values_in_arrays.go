package main

import "fmt"

func main(){

    // this is how you create arrays in Go
    // this creates a string array for holding 3 strings
    var colors [3]string
    colors[0] = "Red"
    colors[1] = "Green"
    colors[2] = "Blue"
    fmt.Println(colors)
    fmt.Println(colors[1])

    // this declare an array and its items in a single statement
    // using an array literal
    var numbers = [5]int{5,3,1,2,4}
    fmt.Println(numbers)

    fmt.Println("Number of colors:", len(colors))
    fmt.Println("Number of numbers:", len(numbers))

}
