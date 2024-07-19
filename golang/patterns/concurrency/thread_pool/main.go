package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

type Worker struct {
	ID          int
	TaskChannel chan Task
	WorkerPool  chan chan Task
	Quit        chan bool
}

func NewWorker(id int, workerPool chan chan Task) Worker {
	return Worker{
		ID:          id,
		TaskChannel: make(chan Task),
		WorkerPool:  workerPool,
		Quit:        make(chan bool),
	}
}

func (w Worker) Start(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			w.WorkerPool <- w.TaskChannel // add the worker's task channel to the worker pool

			select {
			case task := <-w.TaskChannel:
				fmt.Printf("Worker %d processing task %d\n", w.ID, task.ID)
				time.Sleep(1 * time.Second)
			case <-w.Quit:
				fmt.Printf("Worker %d stopping\n", w.ID)
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
	WorkerPool chan chan Task
	MaxWorkers int
	TaskQueue  chan Task
	Workers    []Worker
	wg         *sync.WaitGroup
}

func NewDispatcher(maxWorkers int, taskQueue chan Task) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Task, maxWorkers),
		MaxWorkers: maxWorkers,
		TaskQueue:  taskQueue,
		wg:         &sync.WaitGroup{},
		Workers:    make([]Worker, maxWorkers),
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i+1, d.WorkerPool)
		d.Workers[i] = worker
		d.wg.Add(1)
		worker.Start(d.wg)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case task := <-d.TaskQueue:
			taskChannel := <-d.WorkerPool
			taskChannel <- task
		}
	}
}

func (d *Dispatcher) Stop() {
	for _, worker := range d.Workers {
		worker.Stop()
	}

	d.wg.Wait()
}

func main() {
	const numTasks = 10
	const numWorkers = 3

	taskQueue := make(chan Task, numTasks)

	dispatcher := NewDispatcher(numWorkers, taskQueue)
	dispatcher.Run()

	for i := 1; i <= numTasks; i++ {
		task := Task{ID: i}
		taskQueue <- task
	}

	time.Sleep(5 * time.Second)

	dispatcher.Stop()
	fmt.Println("All workers stopped")

}
