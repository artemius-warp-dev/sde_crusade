package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Pair struct to store character count and character
type Pair struct {
	count int
	ch    rune
}

// MaxHeap to implement a priority queue
type MaxHeap []Pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count } // Max-heap based on count
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push element onto heap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

// Pop element from heap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func longestDiverseString(a, b, c int) string {
	pq := &MaxHeap{}
	heap.Init(pq)

	if a > 0 {
		heap.Push(pq, Pair{a, 'a'})
	}
	if b > 0 {
		heap.Push(pq, Pair{b, 'b'})
	}
	if c > 0 {
		heap.Push(pq, Pair{c, 'c'})
	}

	var result strings.Builder

	for pq.Len() > 0 {
		top := heap.Pop(pq).(Pair)

		if result.Len() >= 2 && result.String()[result.Len()-1] == byte(top.ch) && result.String()[result.Len()-2] == byte(top.ch) {
			if pq.Len() == 0 {
				break
			}

			second := heap.Pop(pq).(Pair)
			result.WriteRune(second.ch)

			if second.count-1 > 0 {
				heap.Push(pq, Pair{second.count - 1, second.ch})
			}

			heap.Push(pq, top)
		} else {
			result.WriteRune(top.ch)

			if top.count-1 > 0 {
				heap.Push(pq, Pair{top.count - 1, top.ch})
			}

		}
	}

	return result.String()

}

// Driver function to test cases
func main() {
	testCases := []struct {
		a, b, c int
	}{
		{1, 1, 7},
		{2, 2, 1},
		{7, 2, 0},
		{0, 0, 0},
		{10, 5, 3},
		{3, 3, 3},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\t a: %d, b: %d, c: %d\n", i+1, tc.a, tc.b, tc.c)
		result := longestDiverseString(tc.a, tc.b, tc.c)
		fmt.Printf("\n\t Longest Happy String: %s\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
