package main

import (
	"fmt"
	"sync"
)

/*
-Please refer to the snippet 30, 31, 32, 33, 34 to understand benefits of using channels
-This example accepts three values
 -cChanSize --> identifies the buffered channel size
 -cWorkerCount --> this defines the number of worker routines that are going to work on the channel
 -cTaskCount --> this defines how many tasks that are going to be carried out
-In this example, taskLoader function will load the tasks that have to be carried out by the workers. When the program
 starts, it initially starts loading the tasks into a buffered channel in a separate go routine. The idea is to start
 workers to consume from the channel as soon as they become available.
-Buffered channels
 A buffered channel is a channel with a capacity to hold. You can think of a channel as a pipe, where a go routine
 checks an incoming stream of items. In the case of buffered channels, there can be multiple items in the channel for
 a given time. If the channel is empty, all the go routines that are looking for incoming items from the channel are
 blocked until there is an item available in the channel.
 Multiple go routines can receive items from a buffered channel one after another without waiting for other go routines
 to finish.
*/

const (
	cChanSize int = 20
	cWorkerCount int = 5
	cTaskCount int = 20000
)

var tasks chan int
var wg sync.WaitGroup

func main() {
	// create counting semaphore to wait until the counting process is finished
	wg.Add(cWorkerCount)

	// this creates a buffered channel with type int and size of chanSize
	tasks = make(chan int, cChanSize)

	go taskLoader()

	for i := 1; i <= cWorkerCount; i++ {
		go worker(i)
	}

	wg.Wait()
}

func taskLoader() {
	for i := 1; i <= cTaskCount; i++ {
		tasks <- i
	}
	close(tasks)
}

func worker(workerNum int) {
	defer wg.Done()
	for {
		task, ok := <- tasks
		if !ok {
			fmt.Printf("worker %v shutting down\n", workerNum)
			break
		}
		fmt.Printf("worker %v processing task num %v\n", workerNum, task)
	}
}