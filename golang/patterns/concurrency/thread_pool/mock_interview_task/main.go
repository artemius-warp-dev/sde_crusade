package main

import (
	"fmt"
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

		defer wg.Done()

		for {
			w.WorkerPool <- w.TaskChanel

			select {
			case task := <-w.TaskChanel:
				fmt.Println(task, " -  ", w.ID)
				time.Sleep(1 * time.Second)
			case <-w.Quit:
				fmt.Println("Quit")
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
		worker := NewWorker(i+1, d.WorkerPool)
		d.Workers[i] = worker
		d.wg.Add(1)
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
	for _, worker := range d.Workers {
		worker.Stop()
	}

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
	}

	dispatcher.Stop()
	//time.Sleep(20 * time.Second)
}
