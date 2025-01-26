package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

// Structure for MaxHeap
type MaxHeap []int

// newMaxHeap function intializes an instance of the MaxHeap
func newMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

// Len function returns the length of the MaxHeap
func (h MaxHeap) Len() int {
	return len(h)
}

// Empty returns true if the MaxHeap is empty, false otherwise
func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

// Greater returns true if the first of the given elements is greater than the second one
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap function swaps the values at the given indices
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() int {
	return h[0]
}

// Push function inserts the element in the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func largestInteger(num int) int {
	digits := make([]int, 0)
	numStr := strconv.Itoa(num)
	for _, ch := range numStr {
		digit, _ := strconv.Atoi(string(ch))
		digits = append(digits, digit)
	}

	oddHeap := newMaxHeap()
	evenHeap := newMaxHeap()

	for _, d := range digits {
		if d%2 == 0 {
			heap.Push(evenHeap, d)
		} else {
			heap.Push(oddHeap, d)
		}
	}

	result := []int{}

	for _, d := range digits {
		if d%2 == 0 {
			result = append(result, heap.Pop(evenHeap).(int))
		} else {
			result = append(result, heap.Pop(oddHeap).(int))
		}
	}

	outputNumberStr := ""
	for _, d := range result {
		outputNumberStr += strconv.Itoa(d)
	}

	outputNumber, _ := strconv.Atoi(outputNumberStr)
	return outputNumber

}

func main() {
	testCases := []int{1234, 65875, 4321, 2468, 98123}
	for _, num := range testCases {
		fmt.Printf("\tInput number: %d\n", num)
		fmt.Printf("\n\tOutput number: %d\n", largestInteger(num))
		fmt.Println(strings.Repeat("-", 100))
	}
}
