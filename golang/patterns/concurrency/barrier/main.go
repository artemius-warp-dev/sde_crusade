package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 5

func worker(id int, wg *sync.WaitGroup, cond *sync.Cond, ready *int) {
	defer wg.Done()

	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d reached the barrier\n", id)

	cond.L.Lock()
	*ready++
	if *ready == numWorkers {
		cond.Broadcast() // Signal all workers that the barrier is reached
	}

	for *ready < numWorkers {
		cond.Wait() // Wait for the signal to proceed
	}

	cond.L.Unlock()

	fmt.Printf("Worker %d proceeding with work\n", id)
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	ready := 0

	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, cond, &ready)
	}

	wg.Wait()
	fmt.Println("All workers have completed their work")

}
