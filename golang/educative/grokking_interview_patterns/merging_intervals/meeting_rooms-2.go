package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *minHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *minHeap) Top() any {
	return (*h)[0]
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findSets(intervals [][]int) int {

	// Replace this placeholder return statement with your code

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &minHeap{}
	heap.Init(h)

	heap.Push(h, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		curStartTime := intervals[i][0]
		fmt.Println(h.Top())
		if curStartTime >= h.Top().(int) {
			heap.Pop(h)
		}
		fmt.Println(curStartTime, *h)
		heap.Push(h, intervals[i][1])

	}

	return h.Len()
}
