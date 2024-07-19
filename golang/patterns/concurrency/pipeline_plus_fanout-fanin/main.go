package main

import (
	"fmt"
	"strings"
	"sync"
)

func worker(id int, tasks <-chan string, results chan<- string, processFunc func(string) string) {
	for task := range tasks {

		fmt.Printf("Worker %d processing task: %s\n", id, task)
		results <- processFunc(task)
	}
}

func stage1(input <-chan string) <-chan string {

	output := make(chan string)
	tasks := make(chan string)

	const numWorkers = 3
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			worker(id, tasks, output, strings.ToUpper)
		}(i)
	}

	go func() {
		for task := range input {
			tasks <- task
		}
		close(tasks)
		wg.Wait()
		close(output)
	}()

	return output
}

func stage2(input <-chan string) <-chan string {

	output := make(chan string)
	tasks := make(chan string)

	const numWorkers = 3

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, tasks, output, reverse)
		}(i)
	}

	go func() {
		for task := range input {
			tasks <- task
		}
		close(tasks)
		wg.Wait()
		close(output)
	}()

	return output
}

func reverse(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input := make(chan string)

	go func() {
		defer close(input)

		input <- "hello"
		input <- "world"
		input <- "gopher"
	}()

	pipeline := stage2(stage1(input))

	for result := range pipeline {
		fmt.Printf("Final result: %s\n", result)
	}
}
