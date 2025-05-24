package main

import (
	"container/heap"
	"fmt"
	"sort"
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

func maxSubsequence(nums []int, k int) []int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for i, num := range nums {
		heap.Push(minHeap, Element{value: num, index: i})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	result := *minHeap

	sort.Slice(result, func(i, j int) bool {
		return result[i].index < result[j].index
	})

	subsequence := make([]int, 0, k)
	for _, el := range result {
		subsequence = append(subsequence, el.value)
	}

	return subsequence
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 4, 3, 3}, 2},
		{[]int{2, 1, 3, 3}, 2},
		{[]int{-1, -2, 3, 4}, 3},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{9, -1, -3, 8, 7}, 2},
	}

	for i, tc := range testCases {
		fmt.Printf("%d\tnums = %v, k = %d\n", i+1, tc.nums, tc.k)
		result := maxSubsequence(tc.nums, tc.k)
		fmt.Printf("\n\tResult: %v\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
