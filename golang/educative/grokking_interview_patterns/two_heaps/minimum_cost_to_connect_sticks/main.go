package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
	totalCost := 0

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, stick := range sticks {
		heap.Push(minHeap, stick)
	}

	for minHeap.Len() > 1 {
		first := heap.Pop(minHeap).(int)
		second := heap.Pop(minHeap).(int)

		cost := first + second
		totalCost += cost

		heap.Push(minHeap, cost)

	}

	return totalCost

}

func main() {
	testCases := [][]int{
		{2, 4, 3},
		{1, 8, 3, 5},
		{5},
		{1, 2, 3, 4, 5},
		{7, 6, 8, 10},
	}

	for i, sticks := range testCases {
		fmt.Printf("%d.\tsticks: %v\n", i+1, sticks)
		fmt.Printf("\tMinimum cost: %d\n", connectSticks(sticks))
		fmt.Println(strings.Repeat("-", 100))
	}
}
