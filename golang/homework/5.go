package main

// Домашнее задание
// Параллельное исполнение
// Написать функцию для параллельного выполнения N заданий (т.е. в N параллельных горутинах).

// Функция принимает на вход:
// - слайс с заданиями `[]func() error`;
// - число заданий которые можно выполнять параллельно (`N`);
// - максимальное число ошибок после которого нужно приостановить обработку.

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func raise_error_after_rand_time() error {
	fmt.Println("Doing...")
	rand.Seed(time.Now().UnixNano())
	randomTime := time.Duration(rand.Intn(5)+1) * time.Second
	time.Sleep(randomTime)

	return errors.New("Error: Unexpected behaviour")
}

func process(tasks []func() error, number int, errors_cap int) {
	errorCh := make(chan interface{}, errors_cap)
	finishCh := make(chan interface{})

	for i := 0; i < number; i++ {
		fmt.Printf("Task %v started \n", i+1)
		go do_process(tasks[i], errorCh, finishCh)
	}

	select {
	case <-finishCh:
		fmt.Println("The max errors is reach")
		return
	}
}

func do_process(task func() error, errorCh chan<- interface{}, finishCh chan<- interface{}) {
	err := task()
	if err != nil {
		errorCh <- "error"
		fmt.Printf("Error has been written. errorCh len: %v \n", len(errorCh))
	}

	if len(errorCh) == cap(errorCh) {
		close(finishCh)
	} else {
		do_process(task, errorCh, finishCh)
	}
}

func main() {
	tasks := []func() error{}
	n := 10
	for i := 0; i < n; i++ {
		tasks = append(tasks, raise_error_after_rand_time)
	}

	process(tasks, n, 20)

}
