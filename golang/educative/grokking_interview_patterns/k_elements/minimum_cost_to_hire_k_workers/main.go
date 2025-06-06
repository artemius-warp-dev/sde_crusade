package main

import (
	"container/heap"
	"fmt"
	"sort"
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

func minCostToHireWorkers(quality []int, wage []int, k int) float64 {
	type Worker struct {
		ratio   float64
		quality int
	}

	workers := make([]Worker, len(quality))
	for i := range quality {
		workers[i] = Worker{
			ratio:   float64(wage[i]) / float64(quality[i]),
			quality: quality[i],
		}
	}

	sort.Slice(workers, func(i, j int) bool {
		return workers[i].ratio < workers[j].ratio
	})

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	totalQuality := 0
	minCost := 1e18

	for _, worker := range workers {

		heap.Push(maxHeap, worker.quality)
		totalQuality += worker.quality

		if maxHeap.Len() > k {
			totalQuality -= heap.Pop(maxHeap).(int)
		}

		if maxHeap.Len() == k {
			cost := worker.ratio * float64(totalQuality)
			if cost < minCost {
				minCost = cost
			}
		}

	}

	return minCost

}

// Driver function to test the solution
func main() {
	qualities := [][]int{
		{10, 20, 5},
		{3, 1, 10, 10, 1},
		{4, 5, 6},
		{2, 3, 1},
		{10, 10, 10},
	}

	wages := [][]int{
		{70, 50, 30},
		{4, 8, 2, 2, 7},
		{8, 10, 12},
		{5, 6, 2},
		{50, 60, 70},
	}

	kValues := []int{2, 3, 2, 2, 2}

	for i := range qualities {
		fmt.Printf("%d.\tqualities: %v\n", i+1, qualities[i])
		fmt.Printf("\twages: %v\n", wages[i])
		fmt.Printf("\tk: %d\n", kValues[i])

		result := minCostToHireWorkers(qualities[i], wages[i], kValues[i])
		fmt.Printf("\n\tMinimum cost to hire %d workers = %.2f\n", kValues[i], result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
