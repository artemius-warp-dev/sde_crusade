package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for n := range tasks {
		fmt.Println(n)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d porcessed %d\n", id, n)
		results <- n * n
	}
}

func main() {

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	// Fan-out
	go func() {
		for i := 1; i <= 5; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// Fan-in
	for i := 1; i <= 5; i++ {
		fmt.Printf("Result: %d\n", <-results)
	}

}
