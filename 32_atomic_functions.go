package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
-Please refer to the snippet number 31 (31_race_conditions.go) before going through this snippet
-Atomic functions
 Atomic functions provide low-level locking mechanisms. Atomic functions lock the variable
 from reading/writing while another goroutine operates on it.
*/

/*
-Following example introduces the usage of atomic functions
-As I described in the snippet number 31, goroutines do not know about each other and they do not
 care about each others' operations. We can overcome this situation by using atomic functions to stop
 the other goroutines from reading/writing while the current one works on it.
-In this case, the incrementCounter function utilizes the atomic.AddInt64 function. This function call
 prevents other goroutines from running any operation on the given memory address while the current
 goroutine works on it.
-In this case, the goroutines do not overwrite on each others' data and it eventually outputs the number
 20000 (10000 + 10000) - That we can not output in the snippet number 31
*/

var counter int64

func main() {

	// create counting semaphores for two goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		incrementCounter(10000)
		fmt.Printf("1st goroutine: %v\n", counter)
	}()

	go func() {
		defer wg.Done()
		incrementCounter(10000)
		fmt.Printf("2nd goroutine: %v\n", counter)
	}()

	wg.Wait()
}

func incrementCounter(limit int) {
	for i := 0; i < limit; i++ {
		// this is an atomic AddInt64 function
		// this function provides the atomicity for the add operation in the given int64 memory address
		atomic.AddInt64(&counter, 1)
	}
}
