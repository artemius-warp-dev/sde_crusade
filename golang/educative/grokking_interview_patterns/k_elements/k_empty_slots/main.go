package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h MinHeap) Min() int {
	if len(h) == 0 {
		return math.MaxInt
	}
	return h[0]
}

func kEmptySlots(bulbs []int, k int) int {
	n := len(bulbs)
	if n == 0 {
		return -1
	}

	// days[i] = the day bulb at position i+1 is turned on
	days := make([]int, n)
	for day, bulb := range bulbs {
		days[bulb-1] = day + 1
	}

	result := math.MaxInt
	h := &MinHeap{}

	// Iterate through positions
	for i := 0; i < n; i++ {
		heap.Push(h, days[i])

		// When we have more than k middle bulbs, remove leftmost
		if i > k {
			// remove days[i-k-1] from heap
			// Note: Go’s heap doesn’t support remove arbitrary efficiently,
			// so for clarity, we rebuild a new heap for each window here.
			*h = (*h)[1:]
			heap.Init(h)
		}

		if i >= k+1 {
			left := i - k - 1
			right := i
			leftDay := days[left]
			rightDay := days[right]
			midMin := math.MaxInt

			// find min in-between bulbs
			for j := left + 1; j < right; j++ {
				if days[j] < midMin {
					midMin = days[j]
				}
			}

			if max(leftDay, rightDay) < midMin {
				result = min(result, max(leftDay, rightDay))
			}
		}
	}

	if result == math.MaxInt {
		return -1
	}
	return result
}

// ---------- Helpers ----------
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	testCases := []struct {
		bulbs []int
		k     int
	}{
		{[]int{1, 3, 2}, 1},
		{[]int{1, 2, 3}, 1},
		{[]int{2, 5, 1, 4, 3}, 1},
		{[]int{3, 1, 5, 4, 2}, 1},
		{[]int{2, 4, 1, 3}, 0},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tbulbs: %v, k: %d\n", i+1, tc.bulbs, tc.k)
		result := kEmptySlots(tc.bulbs, tc.k)
		fmt.Printf("\tEarliest Day: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
