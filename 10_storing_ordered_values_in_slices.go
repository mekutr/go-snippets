package main

import (
	"fmt"
	"sort"
)

/*
When you declare a slice, the run time creates the underlying array for you,
allocates the required memory and then returns the requested slice
*/

func main() {
	// opening and closing bracket without a number means, it is a slice
	// if I were put a number in, that would mean it's an array
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// this will add the "Purple" string to the colors slice
	// and will return a new reference for the new slice
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)]) // here we don't need the len(colors) because it is the default
	// colors = append(colors[1:]) // this is same as above
	fmt.Println(colors)

	colors = append(colors[0 : len(colors)-1]) // here we don't need 0 because it is the default
	// colors = append(colors[:len(colors)-1]) // this is same as above
	fmt.Println(colors)

	// we can also declare a slice with an initial size
	// by using the built-in make function
	// make takes three arguments
	// 1 - the type of the slices items
	// 2 - the initial length
	// 3 - the optional capacity that caps the number of items the slice will contain
	//     if you leave out the capacity, it sets it to the slice's initial size
	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156
	fmt.Println(numbers)

	numbers = append(numbers, 255)
	fmt.Println(numbers)
	fmt.Println(cap(numbers)) // cap function returns the current capacity of a slice

	sort.Ints(numbers) // this will sort the numbers slice in ascending order
	fmt.Println(numbers)
}
