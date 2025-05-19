package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Element stores the number, its prime, and the index of the last used element
type Element struct {
	value int // The next super ugly number
	prime int // The prime factor that generated this number
	index int // The index in the ugly list that was last used for this prime
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

func nthSuperUglyNumber(n int, primes []int) int {
	ugly := []int{1}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, prime := range primes {
		heap.Push(minHeap, Element{value: prime, prime: prime, index: 0})
	}

	for len(ugly) < n {
		top := heap.Pop(minHeap).(Element)
		nextUgly, prime, index := top.value, top.prime, top.index

		if nextUgly != ugly[len(ugly)-1] {
			ugly = append(ugly, nextUgly)
		}

		heap.Push(minHeap, Element{value: prime * ugly[index+1], prime: prime, index: index + 1})
	}

	return ugly[n-1]

}

func main() {
	nValues := []int{12, 1, 15, 10, 8}
	primesList := [][]int{
		{2, 7, 13, 19},
		{2, 3, 5},
		{3, 5, 7},
		{2, 5, 11},
		{3, 11, 17},
	}

	for i := 0; i < len(nValues); i++ {
		fmt.Printf("%d.\tn: %s\n", i+1, fmt.Sprintf("%d", nValues[i]))

		var primeStrs []string
		for _, prime := range primesList[i] {
			primeStrs = append(primeStrs, fmt.Sprintf("%d", prime))
		}
		fmt.Printf("\tprimes: [%s]\n", strings.Join(primeStrs, ", "))

		result := nthSuperUglyNumber(nValues[i], primesList[i])
		fmt.Printf("\n\t%dth super ugly number is %d\n", nValues[i], result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
