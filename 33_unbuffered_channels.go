package main

import (
	"fmt"
	"sync"
)

/*
-Please refer to the snippet 30, 31, 32 and 33 to understand benefits of using channels
-This example accepts two values
 -cLimitPerRoutine --> this identifies the upper limit value that a go routine can increment the value of counter
 -cFinalResult --> this identifies the final result that we want to reach once the main terminates
-In this example, initCounter function will make a decision of how many go routines that it needs to fork by checking
 the limitPerRoutine and finalResult values. For example, if limitPerRoutine is 10000 and finalResult is 20000, it will
 fork 2 go routines to reach the final value.
-Unbuffered channels
 An unbuffered channel is a a channel with no capacity to hold. You can think of a channel as a pipe, where a go routine
 checks an incoming stream of items. In the case of unbuffered channels, there can only be one item in the channel for
 a given time. If the channel is empty, all the go routines that are looking for incoming items from the channel are
 blocked until there is an item available in the channel.
 Since only one go routine operate on a channel in a given time, other go routines still have to wait until the channel
 is available again. For example, if you have 100 go routines waiting an item coming from the channel at the same time,
 only one of them can operate on it once the item becomes available in the channel. The other 99 routines should still
 wait until an item becomes available in the channel again.
*/

var counter chan int64
var wg sync.WaitGroup

const (
	cLimitPerRoutine int64 = 10000
	cFinalResult     int64 = 20000
)

func main() {
	// create counting semaphore to wait until the counting process is finished
	wg.Add(1)

	// this creates an unbuffered channel with type int64
	counter = make(chan int64)

	// this function uses the limitPerRoutine and finalResult values to come up with the
	// total number of go routines that are going to operate to reach the final result
	go initCounter(cLimitPerRoutine, cFinalResult)

	wg.Wait()
}

// this function decides how many go routines have to be created to reach the finalResult value
func initCounter(limitPerRoutine int64, finalResult int64) {
	remainder := finalResult % limitPerRoutine
	routineCount := finalResult / limitPerRoutine

	i := int64(0)
	for ; i < routineCount; i++ {
		go incrementCounter(i+1, limitPerRoutine, finalResult)
	}
	// if there is a remainder, an additional go routine has to be created to reach the finalResult value
	if remainder != 0 {
		go incrementCounter(i+1, remainder, finalResult)
	}

	// this puts the initial value of 0 into the channel to start the process
	// if the channel stays empty and we do not put any value in it, none of the
	// go routines can operate.
	counter <- 0
}

// this function takes a value from the channel, operates on it, and puts
// the value that is updated back to the channel again.
func incrementCounter(sequence int64, limit int64, finalResult int64) {
	num := <-counter
	for i := int64(0); i < limit; i++ {
		num++
	}
	fmt.Printf("goroutine %v : %v\n", sequence, num)

	// if the final result has not reached yet, there must be other
	// go routines waiting to operate on the value that is going to be
	// inserted into the channel
	if num != finalResult {
		counter <- num
	} else {
		// if the final result is reached, close the channel and signal the wait group
		go endCounter()
	}
}

func endCounter() {
	close(counter)
	wg.Done()
}
