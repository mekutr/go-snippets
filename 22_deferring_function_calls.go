package main

import (
    "fmt"
)

/*
    -Go's "defer" statement does exactly what it says. It defers execution of a block of
     code until everything else in the current function is finished

    -Each time you call the "defer" statement, it adds an instruction to a stack and when the
     deferred statements are executed, they' re handled in last in first out order, known as LIFO
*/

func main() {
    // defer statement runs "fmt.Println("Close the file!")" after everything else
    // is finished in this function -- because this is the first one deferred
    defer fmt.Println("Close the file!")
    fmt.Println("Open the file!")

    defer fmt.Println("Statement 1")
    defer fmt.Println("Statement 2")

    // the defer keyword only works within a function.
    // any deferred statements will be executed at the end of the current function,
    // rather than waiting for the entire application
    myFunc()

    defer fmt.Println("Statement 3")
    defer fmt.Println("Statement 4")
    fmt.Println("Undeferred statment")

    // any values that you refer to in a deferred statement
    // is evaluated at the moment that the statement is deferred
    // rather than when it's executed

    x := 1000
    defer fmt.Println("Value of x:", x) // this will print 1000 as the value of x because of the reason above
    x++
    fmt.Println("Value of x after incrementing", x)
}

func myFunc() {
    defer fmt.Println("Deffered in the function")
    fmt.Println("Not deferred in the function")
}