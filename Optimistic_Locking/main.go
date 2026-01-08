package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32 = 0

func incrementCount(wg *sync.WaitGroup) {
	defer wg.Done()
	oldValue := atomic.LoadInt32(&count)
	newValue := oldValue + 1

	if !atomic.CompareAndSwapInt32(&count, oldValue, newValue) {
		fmt.Println("Failed to increment count")
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go incrementCount(&wg)
	go incrementCount(&wg)
	wg.Wait()
	fmt.Println("Count:", count)
}
