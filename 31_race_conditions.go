package main

import (
	"fmt"
	"sync"
)

/*
-When multiple goroutines are not synchronized while reading/writing from/to a shared resource,
 they might attempt to read and write to that resource without waiting each others' turn. This
 condition is called a race condition.
-Goroutines do not respect the atomic transactions that the other goroutines has to handle if
 the goroutines are not designed in that respect. If the atomic transactions that each goroutine
 has to handle is not identified while designing the software, it introduces unwanted bugs into
 the software.
*/

/*
-Following example introduces a race condition between two goroutines
-In this case, second goroutine will not respect the first goroutine and it will try to increment
 the counter as soon as it takes its turn. While the expected output is 10000 after the first goroutine
 is finished executing and 20000 after the second goroutine is finished executing, it will always create
 different results after the following is executed.
-The reason is because there is no atomicity controller between two goroutines to make the second one
 wait until the first one finishes.
*/

var counter int

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
		counter++
	}
}
