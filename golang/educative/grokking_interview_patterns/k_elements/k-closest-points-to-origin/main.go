package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Structure for MaxHeap of pairs (represented as slices of integers)
type MaxHeap [][]int

// newMaxHeap function initializes an instance of the MaxHeap
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

// Less returns true if the first of the given elements is greater than the second one
// Here, we compare the first integers, and if they're equal, we compare the second integers
func (h MaxHeap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] > h[j][1]
	}
	return h[i][0] > h[j][0]
}

// Swap function swaps the values at the given indices
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() []int {
	return h[0]
}

// Push function inserts the element in the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func squaredDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func kClosest(points [][]int, k int) [][]int {
	closest := newMaxHeap()

	for i := 0; i < k; i++ {
		dist := squaredDistance(points[i])
		heap.Push(closest, []int{dist, i})
	}

	for i := k; i < len(points); i++ {
		dist := squaredDistance(points[i])

		if dist < closest.Top()[0] {
			heap.Pop(closest)
			heap.Push(closest, []int{dist, i})
		}
	}

	result := [][]int{}
	for closest.Len() > 0 {
		idx := heap.Pop(closest).([]int)[1]
		result = append(result, points[idx])
	}

	return result

}

// Driver code
func main() {
	pointsArr := [][][]int{
		{{1, 3}, {3, 4}, {2, -1}},
		{{1, 3}, {2, 4}, {2, -1}, {-2, 2}, {5, 3}, {3, -2}},
		{{1, 3}, {5, 3}, {3, -2}, {-2, 2}},
		{{2, -1}, {-2, 2}, {1, 3}, {2, 4}},
		{{1, 3}, {2, 4}, {2, -1}, {-2, 2}, {5, 3}, {3, -2}, {5, 3}, {3, -2}},
	}

	kArr := []int{2, 3, 1, 4, 5}

	for i := 0; i < len(pointsArr); i++ {
		fmt.Printf("%d.\tpoints: [", i+1)
		for j := 0; j < len(pointsArr[i]); j++ {
			fmt.Printf("[%d, %d]", pointsArr[i][j][0], pointsArr[i][j][1])
			if j < len(pointsArr[i])-1 {
				fmt.Print(", ")
			}
		}
		fmt.Printf("]\n")
		fmt.Printf("\tk: %d\n\n", kArr[i])

		result := kClosest(pointsArr[i], kArr[i])

		fmt.Printf("\t%d closest points to origin: [", kArr[i])
		for j := 0; j < len(result); j++ {
			fmt.Printf("[%d, %d]", result[j][0], result[j][1])
			if j < len(result)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Printf("]\n")
		fmt.Println(strings.Repeat("-", 100))
	}
}
