package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100000000
var CONCURRENCY = 10
var totalPrimeNumbers int32 = 0

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	// Use atomic operation to safely increment the counter
	// atomice means here that the operation is done in a thread-safe manner which means that multiple goroutines can safely increment the counter
	// in simple words not in programming words it means that the operation is done in a way that it cannot be interrupted by other operations
	atomic.AddInt32(&totalPrimeNumbers, 1)
}


// doBatch checks for prime numbers in the range [nstart, nend)
// and signals completion via the WaitGroup
// it would basically check for prime numbers in the given range and print the time taken to complete the batch
func doBatch(name string, wg *sync.WaitGroup, nstart int, nend int) {
	defer wg.Done()
	start := time.Now()
	for i := nstart; i < nend; i++ {
		checkPrime(i)
	}
	fmt.Printf("thread %s [%d, %d] completed in %s\n", name, nstart,
		nend, time.Since(start))
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	nstart := 3
	batchSize := int(float64(MAX_INT) / float64(CONCURRENCY))

	// here we are creating 10 threads to check for prime numbers concurrently
	// each thread will check a batch of numbers for our case its 10 million numbers per thread
	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)
		go doBatch(strconv.Itoa(i), &wg, nstart, nstart+batchSize)
		nstart += batchSize
	}
	wg.Add(1)
	go doBatch(strconv.Itoa(CONCURRENCY-1), &wg, nstart, MAX_INT)
	wg.Wait()

	fmt.Println("Checking till", MAX_INT, "found", totalPrimeNumbers+1,
		"prime numbers. took", time.Since(start))
}
