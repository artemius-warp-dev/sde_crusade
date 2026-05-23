package main

import (
	"container/heap"
	"fmt"
	"math"
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

func maxScore(nums []int, k int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, num := range nums {
		heap.Push(maxHeap, num)
	}

	score := 0

	for i := 0; i < k; i++ {
		largest := heap.Pop(maxHeap).(int)
		score += largest

		reduced := int(math.Ceil(float64(largest) / 3))
		heap.Push(maxHeap, reduced)
	}

	return score
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{12, 18, 24, 6}, 3},
		{[]int{7, 14, 3}, 2},
		{[]int{50, 20, 15, 10, 5}, 4},
		{[]int{8, 16, 5, 12, 3, 7}, 3},
		{[]int{50, 50, 50, 50, 50, 50, 50, 50}, 5},
	}

	for i, testCase := range testCases {
		fmt.Printf("%d.\tnums: %v\n", i+1, testCase.nums)
		fmt.Printf("\tk: %d\n", testCase.k)
		result := maxScore(testCase.nums, testCase.k)
		fmt.Printf("\n\tMaximum Score: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
