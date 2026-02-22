package main

import (
	"container/heap"
	"fmt"
)

// A MinHeap implements heap.Interface and holds ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // min-heap
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

func findLeastNumOfUniqueInts(arr []int, k int) int {
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}

	h := &MinHeap{}
	heap.Init(h)
	for _, freq := range freqMap {
		heap.Push(h, freq)
	}

	counter := 0

	for h.Len() > 0 {
		counter += (*h)[0]

		if counter > k {
			return h.Len()
		}
		heap.Pop(h)
	}
	return 0
}

func main() {
	arrays := [][]int{
		{4, 3, 1, 1, 3, 3, 2},
		{5, 5, 4},
		{2, 4, 1, 8, 3, 5, 1, 3},
		{1, 1, 1, 2, 2, 3, 3, 3, 3},
		{5, 5, 5, 5, 6, 7, 8},
	}
	ks := []int{3, 1, 3, 4, 2}

	for i := 0; i < len(arrays); i++ {
		arr := arrays[i]
		k := ks[i]

		fmt.Printf("%d.\tarr = [", i+1)
		for j := 0; j < len(arr); j++ {
			if j < len(arr)-1 {
				fmt.Printf("%d, ", arr[j])
			} else {
				fmt.Printf("%d", arr[j])
			}
		}
		fmt.Printf("], k = %d\n", k)

		result := findLeastNumOfUniqueInts(arr, k)
		fmt.Printf("\tLeast Number of Unique Integers: %d\n", result)
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
}
