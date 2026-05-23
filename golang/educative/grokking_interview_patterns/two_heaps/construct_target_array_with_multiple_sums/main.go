package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// MaxHeap struct to define a max-heap
type MaxHeap []int

// Implementing heap.Interface methods for MaxHeap

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Greater than to make it a max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func isPossible(target []int) bool {
	totalSum := 0
	for _, num := range target {
		totalSum += num
	}

	h := &MaxHeap{}
	heap.Init(h)
	for _, num := range target {
		heap.Push(h, num)
	}

	for {
		currentMax := heap.Pop(h).(int)
		remainingSum := totalSum - currentMax

		if currentMax == 1 || remainingSum == 1 {
			return true
		}

		if remainingSum == 0 || currentMax < remainingSum || currentMax%remainingSum == 0 {
			return false

		}

		updatedValue := currentMax % remainingSum
		totalSum = remainingSum + updatedValue

		heap.Push(h, updatedValue)
	}
}

func main() {
	testCases := [][]int{
		{9, 3, 5},
		{1, 1, 1, 2},
		{8, 5},
		{1, 100000},
		{2},
	}

	for i, testCase := range testCases {
		fmt.Printf("%d.\ttarget: [", i+1)
		// Print the target array like Python style: [1, 2, 3]
		for j, val := range testCase {
			fmt.Print(val)
			if j != len(testCase)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")

		// Call IsPossible and print the result
		result := isPossible(testCase)
		fmt.Printf("\toutput: %t\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
