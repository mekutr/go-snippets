package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	// Intn return an integer value starting from 0 up to the ceiling
	// that I provide
	dow := rand.Intn(6) + 1 // since I'll look for days of a week, I added 1 here
	fmt.Println("Day", dow)

	// In Go, code flow automatically jumps past additional
	// cases after finding a match. So you don't need a bunch of
	// redundant break statements

	result := ""

	switch dow {
	case 1:
		result = "It's Sunday!"
	case 7:
		result = "It's Saturday!"
	default:
		result = "It's a weekday!"
	}

	fmt.Println("Day", dow, ",", result)

	result2 := ""
	// you can include an initial statement as part of the if declaration -- x := -42 is an initial statement
	// we can also define the switch statement with conditional expressions
	switch x := -42; {
	case x < 0:
		result2 = "Less than zero"
		//fallthrough // once fallthrough is used switch statement follows executing the next case statement
	case x == 0:
		result2 = "Equal to zero"
	case x > 0:
		result2 = "Greater than zero"
	}

	fmt.Println(result2)
}
