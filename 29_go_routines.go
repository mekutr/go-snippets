package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// runtime.NumCPU() returns how many physical cores your CPU has
	fmt.Printf("Your computer has %v cores.\n", runtime.NumCPU())

	// decide how many logical processors has to be assigned to be used by this executable
	// -- if this line is never declared, go scheduler creates logical processors which is equal to the amount physical cores
	runtime.GOMAXPROCS(1)

	// wait-group is used to create counting semaphores
	// Since all the threads get automatically terminated once the main thread terminates,
	// we have to have a way to wait until all the threads are processed.
	// We can handle this by creating a WaitGroup type by adding the number of go-routines
	// it needs to wait for
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		// Done function call indicates that the go routine is done processing
		// that's why it is a deferred function
		defer wg.Done()
		findPrime("A", 10000)
	}()

	go func() {
		// Done function call indicates that the go routine is done processing
		// that's why it is a deferred function
		defer wg.Done()
		findPrime("B", 10000)
	}()

	// this continues once the equal amount of Done function calls are called.
	wg.Wait()
}

// this function is used for finding the prime numbers
// since this runs in O(n^2), this will take longer time as the input size gets larger
func findPrime(label string, limit int) {
	for outer := 2; outer <= limit; outer++ {
		primeFound := true
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				primeFound = false
			}
		}
		if primeFound {
			fmt.Printf("%v: %v\n", label, outer)
		}
	}
}
