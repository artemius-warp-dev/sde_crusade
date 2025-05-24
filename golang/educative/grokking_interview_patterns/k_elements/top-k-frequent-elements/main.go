package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Pair struct {
	first, second int
}

// MinHeap structure initialization
type MinHeap []Pair

// newMinHeap function initializes an instance of MinHeap
func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

// Len function returns the length of min heap
func (h MinHeap) Len() int {
	return len(h)
}

// Empty function returns true if empty, false otherwise
func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of the MinHeap given their indexes
func (h MinHeap) Less(i, j int) bool {
	return (h[i].second < h[j].second)
}

// Swap function swaps the values of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() Pair {
	return h[0]
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	numFrequencyMap := make(map[int]int)
	for _, n := range nums {
		numFrequencyMap[n] = numFrequencyMap[n] + 1
	}

	topKElements := newMinHeap()

	for num, frequency := range numFrequencyMap {
		heap.Push(topKElements, Pair{num, frequency})
		if topKElements.Len() > k {
			heap.Pop(topKElements)
		}
	}

	topNumbers := make([]int, 0)
	for !topKElements.Empty() {
		pair := topKElements.Top()
		heap.Pop(topKElements)
		topNumbers = append(topNumbers, pair.first)
	}

	return topNumbers

}

// Driver code
func main() {
	nums := [][]int{
		{1, 3, 5, 12, 11, 12, 11, 12, 5},
		{1, 3, 5, 14, 18, 14, 5},
		{2, 3, 4, 5, 6, 7, 7},
		{9, 8, 7, 6, 6, 5, 4, 3, 2, 1},
		{2, 4, 3, 2, 3, 4, 5, 4, 4, 4},
		{1, 1, 1, 1, 1, 1},
		{2, 3},
	}
	k := []int{3, 2, 1, 1, 3, 1, 3}

	for i, num := range nums {
		fmt.Printf("%d.\tInput: (%s, %d)\n", i+1, strings.Replace(fmt.Sprint(num), " ", ", ", -1), k[i])
		fmt.Printf("\tTop %d frequent elements: %s\n", k[i], strings.Replace(fmt.Sprint(topKFrequent(num, k[i])), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
