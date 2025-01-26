package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Heap structure intialization
type Heap [][]int

// newHeap function initializes an instance of the heap
func newHeap() *Heap {
	min := &Heap{}
	heap.Init(min)
	return min
}

// Len function returns the length of the heap
func (h Heap) Len() int {
	return len(h)
}

// Empty function returns true if empty, false otherwise
func (h Heap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of the heap given their indexes
func (h Heap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

// Swap function swaps the values of the elements whose indices are given
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the heap
func (h Heap) Top() interface{} {
	return h[0]
}

// Push function pushes an element into the heap
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

// Pop function pops the element at the top of the heap
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) Heapify(array [][]int) {
	for _, arr := range array {
		heap.Push(h, arr)
	}
}

func tasks(tasks [][]int) int {
	optimalMachines := 0

	machinesAvailable := newHeap()
	tasksList := newHeap()

	tasksList.Heapify(tasks)
	machineInUse := make([]int, 0)

	for tasksList.Empty() == false {
		task := heap.Pop(tasksList).([]int)

		if machinesAvailable.Empty() == false && machinesAvailable.Top().([]int)[0] <= task[0] {
			machineInUse = heap.Pop(machinesAvailable).([]int)

			machineInUse = []int{task[1], machineInUse[1]}

		} else {

			optimalMachines++
			machineInUse = []int{task[1], optimalMachines}
		}
		heap.Push(machinesAvailable, machineInUse)
	}

	return optimalMachines

}

func main() {
	tasksList := [][][]int{
		{
			{1, 1},
			{5, 5},
			{8, 8},
			{4, 4},
			{6, 6},
			{10, 10},
			{7, 7},
		},
		{
			{1, 7},
			{1, 7},
			{1, 7},
			{1, 7},
			{1, 7},
			{1, 7},
		},
		{
			{1, 7},
			{8, 13},
			{5, 6},
			{10, 14},
			{6, 7},
		},
		{
			{1, 3},
			{3, 5},
			{5, 9},
			{9, 12},
			{12, 13},
			{13, 16},
			{16, 17},
		},
		{
			{12, 13},
			{13, 15},
			{17, 20},
			{13, 14},
			{19, 21},
			{18, 20},
		},
	}

	for i, task := range tasksList {
		fmt.Printf("%d.\tTasks = %s\n", i+1, strings.Replace(fmt.Sprint(task), " ", ", ", -1))
		fmt.Printf("\tOptimal number of machines = %d\n", tasks(task))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
