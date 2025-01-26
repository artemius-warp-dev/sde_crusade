package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Set struct {
	n1 int
	n2 int
}

// MinHeap structure intialization
type MinHeap []Set

// newMinHeap function initializes an instance of MinHeap
func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

// Len function returns the length of MinHeap
func (h MinHeap) Len() int {
	return len(h)
}

// Empty function returns true if empty, false otherwise
func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of MinHeap given their indices
func (h MinHeap) Less(i, j int) bool {
	return h[i].n1 < h[j].n1
}

// Swap function swaps the value of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() interface{} {
	return h[0]
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap structure intialization
type MaxHeap []Set

// newMaxHeap function initializes an instance of MaxHeap
func newMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

// Len function returns the length of MaxHeap
func (h MaxHeap) Len() int {
	return len(h)
}

// Empty function returns true if the MaxHeap empty, false otherwise
func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of MaxHeap given their indices
func (h MaxHeap) Less(i, j int) bool {
	return h[i].n1 > h[j].n1
}

// Swap function swaps the value of the elements whose indices are given
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() interface{} {
	return h[0]
}

// Push function pushes an element into the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumCapital(c int, k int, capitals []int, profits []int) int {
	capitalsMinHeap := newMinHeap()
	currentCapital := c
	profitsMaxHeap := newMaxHeap()

	for index, value := range capitals {
		heap.Push(capitalsMinHeap, Set{n1: value, n2: index})

	}

	for i := 0; i < k; i++ {
		for !capitalsMinHeap.Empty() && capitalsMinHeap.Top().(Set).n1 <= currentCapital {
			set := heap.Pop(capitalsMinHeap).(Set)
			u := set.n2
			heap.Push(profitsMaxHeap, Set{n1: profits[u]})
		}

		if profitsMaxHeap.Empty() {
			break
		}

		j := heap.Pop(profitsMaxHeap).(Set).n1
		currentCapital = currentCapital + j
	}

	return currentCapital

}

// Driver code
func main() {
	c := []int{1, 2, 1, 7, 2}
	k := []int{2, 3, 3, 2, 4}
	capitals := [][]int{
		{1, 2, 2, 3},
		{1, 3, 4, 5, 6},
		{1, 2, 3, 4},
		{6, 7, 8, 10},
		{2, 3, 5, 6, 8, 12},
	}
	profits := [][]int{
		{2, 4, 6, 8},
		{1, 2, 3, 4, 5},
		{1, 3, 5, 7},
		{4, 8, 12, 14},
		{1, 2, 5, 6, 8, 9},
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("%d.\tProject capital requirements: %s\n", i+1, strings.Replace(fmt.Sprint(capitals[i]), " ", ", ", -1))
		fmt.Printf("\tProject expected profits: %s\n", strings.Replace(fmt.Sprint(profits[i]), " ", ", ", -1))
		fmt.Printf("\tNumber of projects: %d\n", k[i])
		fmt.Printf("\tStart-up capital: %d\n", c[i])
		fmt.Printf("\tProcessing...\n")
		fmt.Printf("\n\tMaximum capital earned: %d\n", maximumCapital(c[i], k[i], capitals[i], profits[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
