package main

import (
    "fmt"
    "math"
    "math/big" // this is how we can add nested packages
)

func main() {
    //implicitConvert()
    //bigSum()
    circleRadius()
}

func implicitConvert() {

    // Numeric types in Go doesn't implicitly convert numeric types
    // Example: You can't add an int to a float without explicit conversion

    var int1 int = 5
    var float1 float64 = 42
    sum := float64(int1) + float1
    fmt.Printf("Sum: %v, Type: %T", sum, sum)

}

func bigSum() {

    // declares 3 variables of the same type and setting
    // all of their values in a single statement
    i1, i2, i3 := 12, 45, 68
    intSum := i1 + i2 + i3
    fmt.Println("Integer sum:", intSum)

    f1, f2, f3 := 23.5, 65.1, 76.3
    floatSum := f1 + f2 + f3
    // with floating point numbers you might not get the precision you hope for
    fmt.Println("Float sum:", floatSum)

    // use big number types to execute math operations with precision
    var b1, b2, b3, bigSum big.Float
    b1.SetFloat64(23.5)
    b2.SetFloat64(65.1)
    b3.SetFloat64(76.3)

    // & operator is the address of operator
    // it takes the address of b1 and b2 and passes to the Add function
    // the first add call takes the total of b1 and b2 and assigns it to the bigSum
    // the second add call takes the total of bigSum and b3 and assigns it to the bigSum again
    bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
    fmt.Printf("BigSum = %g\n", &bigSum) // ?? for some reason arbitrary-precision arithmetic cannot handle this
}

// The math package also has a whole bunch of functions
// and constants that can help you with your mathematical
// operations
func circleRadius() {

    circleRadius := 15.5
    circumference := circleRadius * math.Pi
    fmt.Printf("Circumference: %.2f\n", circumference)

}
