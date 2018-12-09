package main

import "fmt"

/*
You can loop through your code in Go using the for statement.
Unlike C and Java, there is no while statement in Go but you can
do the same sort of iterating of your code with extended versions of for
*/

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	// traditional for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// traditional for loop
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// this sets the current value of i to the current index
	// as you loop through the slice
	for i := range colors {
		fmt.Println(colors[i])
	}

	// Go doesn't have the while loop but we can accomplish the
	// same thing by using the for loop
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
	}

	// break and continue work the same as other C style languages
	sum = 1
	for true {
		if sum%2 == 0 {
			sum += 11
			continue
		}
		if sum > 1000 {
			break
		}
		sum += sum
	}
	fmt.Println("Sum:", sum)
}
