package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var count int = 0

func incCount() {
	mu.Lock()
	count++
	mu.Unlock()
	// it will decrement the value by 1 after the execution
	wg.Done()
}
func doCount() {
	for i := 0; i < 1000000; i++ {
		// it will increment that value by that number in atomic way
		// here it will increment the value by 1
		wg.Add(1)
		// here we are creating 1000000 threads
		go incCount()
	}
}

func main() {
	count = 0
	doCount()
	wg.Wait() // it will wait until all the threads are completed so its a blocking call
	fmt.Println(count)
}

