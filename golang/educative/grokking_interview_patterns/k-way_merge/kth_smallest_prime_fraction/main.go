package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type MinHeap [][]float64

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]float64))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {

	n := len(arr)
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for j := 1; j < n; j++ {
		heap.Push(minHeap, []float64{float64(arr[0]) / float64(arr[j]), 0, float64(j)})
	}

	for i := 0; i < k-1; i++ {
		top := heap.Pop(minHeap).([]float64)
		numIdx := int(top[1])
		denIdx := int(top[2])

		if numIdx+1 < denIdx {
			heap.Push(minHeap, []float64{float64(arr[numIdx+1]) / float64(arr[denIdx]), float64(numIdx + 1), float64(denIdx)})
		}
	}

	top := heap.Pop(minHeap).([]float64)
	return []int{arr[int(top[1])], arr[int(top[2])]}

}

func main() {
	testCases := []struct {
		arr []int
		k   int
	}{
		{[]int{1, 3, 5, 7, 9, 11}, 2},
		{[]int{1, 7, 23, 29, 47}, 3},
		{[]int{1, 2, 3, 5}, 3},
		{[]int{1, 2, 3, 5}, 1},
		{[]int{1, 13, 17, 19, 23, 29, 31}, 4},
	}

	for idx, testCase := range testCases {
		fmt.Printf("%d.\tArray: %v, k: %d\n", idx+1, testCase.arr, testCase.k)
		result := kthSmallestPrimeFraction(testCase.arr, testCase.k)
		fmt.Printf("\tKth smallest prime fraction is: [%d, %d]\n", result[0], result[1])
		fmt.Println(strings.Repeat("-", 100))
	}
}
