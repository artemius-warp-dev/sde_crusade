package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var urls []string = []string{
	"http://example.com/file1",
	"http://example.com/file2",
	"http://example.com/file3",
	"http://example.com/file4",
	"http://example.com/file5",
}

type Worker struct {
	ID         int
	TaskChanel chan string
	WorkerPool chan chan string
	Quit       chan bool
}

func NewWorker(id int, workerPool chan chan string) Worker {
	return Worker{
		ID:         id,
		TaskChanel: make(chan string),
		WorkerPool: workerPool,
		Quit:       make(chan bool),
	}
}

func (w Worker) Start(wg *sync.WaitGroup) {
	go func() {

		for {
			w.WorkerPool <- w.TaskChanel

			select {
			case task := <-w.TaskChanel:
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				done := make(chan bool)
				go func() {
					n := rand.IntN(3) + 1
					fmt.Printf("Downloaded file from: %s by %d and estimated time is %d second \n", task, w.ID, n)
					time.Sleep(time.Duration(n) * time.Second)
					done <- true
				}()

				select {
				case <-done:
					fmt.Println("Completed:", task)
				case <-ctx.Done():
					fmt.Println("Cancelled:", task)
				}
				wg.Done()

			case <-w.Quit:
				fmt.Println("Worker", w.ID, "quitting")
				return

			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan string
	MaxWorkers int
	TaskQueue  chan string
	Workers    []Worker
	wg         *sync.WaitGroup
}

func NewDispatcher(maxWorkers int, taskQueue chan string) Dispatcher {
	return Dispatcher{
		WorkerPool: make(chan chan string, maxWorkers),
		MaxWorkers: maxWorkers,
		TaskQueue:  taskQueue,
		wg:         &sync.WaitGroup{},
		Workers:    make([]Worker, maxWorkers),
	}
}

func (d Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		d.Workers[i] = worker
		worker.Start(d.wg)
	}
	go d.dispatch()
}

func (d Dispatcher) dispatch() {
	for {
		select {
		case task := <-d.TaskQueue:
			taskChannel := <-d.WorkerPool
			taskChannel <- task
		}
	}
}

func (d Dispatcher) Stop() {
	d.wg.Wait()
}

func main() {
	numTasks := len(urls)
	numWorkers := 3

	taskQueue := make(chan string, numTasks)

	dispatcher := NewDispatcher(numWorkers, taskQueue)
	dispatcher.Run()

	for i := 0; i < numTasks; i++ {
		fmt.Println("Task started:", i)
		taskQueue <- urls[i]
		dispatcher.wg.Add(1)
	}

	dispatcher.Stop()
}
