package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Element stores the number and its original index
type Element struct {
	value int
	index int
}

// MinHeap is a priority queue implemented as a min-heap
type MinHeap []Element

// Implement heap.Interface
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value } // Min-heap based on value
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findRightInterval(intervals [][]int) []int {
	result := make([]int, len(intervals))
	for i := range result {
		result[i] = -1
	}

	startHeap := &MinHeap{}
	endHeap := &MinHeap{}
	heap.Init(startHeap)
	heap.Init(endHeap)

	for i, interval := range intervals {
		heap.Push(startHeap, Element{value: interval[0], index: i})
		heap.Push(endHeap, Element{value: interval[1], index: i})
	}

	for endHeap.Len() > 0 {
		endElem := heap.Pop(endHeap).(Element)

		for startHeap.Len() > 0 && (*startHeap)[0].value < endElem.value {
			heap.Pop(startHeap)
		}

		if startHeap.Len() > 0 {
			result[endElem.index] = (*startHeap)[0].index
		}

	}

	return result

}

func main() {
	testCases := [][][]int{
		{{1, 2}},                         // Single interval, no right interval exists
		{{3, 4}, {2, 3}, {1, 2}},         // Mixed intervals
		{{1, 4}, {2, 3}, {3, 4}},         // Overlapping intervals
		{{5, 6}, {1, 2}, {3, 4}},         // No right intervals for some
		{{1, 3}, {2, 4}, {3, 5}, {4, 6}}, // Sequential intervals
	}

	// Run each test case in a loop
	for i, testCase := range testCases {
		fmt.Printf("%d\tintervals: %v\n", i+1, testCase)
		result := findRightInterval(testCase)
		fmt.Printf("\n\tOutput: %v\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
