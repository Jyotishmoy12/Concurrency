package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type concurrentQueue struct {
	queue []int32
	mu    sync.Mutex
}

func (q *concurrentQueue) Enqueue(item int32) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, item) // This operation is not atomic in nature
}

func (q *concurrentQueue) Dequeue() int32 {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.queue) == 0 {
		panic("Can't deque from an empty queue..")
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *concurrentQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.queue)
}

var wgE sync.WaitGroup
var wgD sync.WaitGroup

func main() {
	q1 := concurrentQueue{
		queue: make([]int32, 0),
	}

	// Enqueue 1 Million items
	for i := 0; i < 1000000; i++ {
		wgE.Add(1)
		go func() {
			q1.Enqueue(rand.Int31())
			wgE.Done()
		}()
	}
	wgE.Wait()
	fmt.Println("After Enqueue", q1.Size())
	for i := 0; i < 1000000; i++ {
		wgD.Add(1)
		go func() {
			q1.Dequeue()
			wgD.Done()
		}()
	}
	wgD.Wait()
	fmt.Println("After Dequeue", q1.Size())
}
