package main

import (
	"fmt"
	"sync"
)

/*
-Please refer to the snippet number 31 (31_race_conditions.go) before going through this snippet.
-This example reaches the same output result (2000) as in snippet number 32 by using a mutex lock
 instead of using an atomic function.
-In this example, the first go-routine will increment the number until 1000 and the second one will
 increment it from 1000 to 2000. This is expected because the for-loop is inside the critical section
 and other goroutines can not reach the critical section until it is released.
-Mutexes has to be preferred if the programmer has to define a critical section instead of introducing
 a lock on a specific variable as I did in snippet number 32.
-Mutexes
 Another way to protect a critical section in your code is by using the mutexes. Mutexes can be
 used to create a critical section in a code and only one goroutine can access the critical section
 after acquiring the mutex lock. If other goroutines try to acquire the lock while another goroutine
 operates on the critical section, their request gets declined.
*/

var counter int64
var mutex sync.Mutex

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
	/*
		once the mutex is acquired by a goroutine, another goroutines
		can not acquire the same lock until it is released by the acquirer.
	*/
	// once this is called other goroutines can not execute the
	// critical section (the section between mutex.Lock and mutex.Unlock)
	mutex.Lock()
	for i := 0; i < limit; i++ {
		counter++
	}
	mutex.Unlock()
}
