package main

import (
	"fmt"
	"sync"
	"time"
)

type Future struct {
	resultChan chan interface{}
	once       sync.Once
}

func NewFuture() *Future {
	return &Future{
		resultChan: make(chan interface{}),
	}
}

func (f *Future) Get() interface{} {
	return <-f.resultChan
}

type Promise struct {
	future *Future
}

func NewPromise() *Promise {
	return &Promise{
		future: NewFuture(),
	}
}

func (p *Promise) Future() *Future {
	return p.future
}

func (p *Promise) SetResult(result interface{}) {
	p.future.once.Do(func() {
		p.future.resultChan <- result
	})

}

func longRunningComputation(promise *Promise, input int) {

	time.Sleep(20 * time.Second)
	result := input * input
	promise.SetResult(result)

}

func main() {

	promise := NewPromise()
	future := promise.Future()

	go longRunningComputation(promise, 50)

	fmt.Println("Doing other work...")

	result := future.Get()

	fmt.Printf("Result of computation: %v\n", result)

}
