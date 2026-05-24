package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Row represents a matrix row with its strength and index.
type Row struct {
	strength int
	index    int
}

// MaxHeap implements a max-heap for Rows.
type MaxHeap []Row

func (h MaxHeap) Len() int { return len(h) }

// Less defines the heap order: higher strength first; if equal, higher index first.
func (h MaxHeap) Less(i, j int) bool {
	if h[i].strength == h[j].strength {
		return h[i].index > h[j].index
	}
	return h[i].strength > h[j].strength
}

func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Row)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func findKWeakestRows(matrix [][]int, k int) []int {
	n := len(matrix[0])

	binarySearch := func(row []int) int {
		low, high := 0, n
		for low < high {
			mid := low + (high-low)/2
			if row[mid] == 1 {
				low = mid + 1
			} else {
				high = mid
			}
		}
		return low
	}

	h := &MaxHeap{}
	heap.Init(h)

	for i, row := range matrix {
		strength := binarySearch(row)
		heap.Push(h, Row{strength, i})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Row).index
	}

	return result
}

// Custom function to print the matrix with commas
func printMatrix(matrix [][]int) {
	fmt.Print("[")
	for i, row := range matrix {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print("[")
		for j, val := range row {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Print(val)
		}
		fmt.Print("]")
	}
	fmt.Print("]")
}

// Helper function to format indexes for printing
func formatIndexes(indexes []int) string {
	var formatted []string
	for _, idx := range indexes {
		formatted = append(formatted, fmt.Sprintf("%d", idx))
	}
	return strings.Join(formatted, ", ")
}

func main() {
	matrixList := [][][]int{
		{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}},
		{{1, 1, 0, 0}, {1, 0, 0, 0}, {1, 1, 1, 1}, {1, 1, 0, 0}},
		{{1, 1}, {1, 1}, {0, 0}, {1, 0}, {1, 1}},
		{{1, 0, 0, 0}, {1, 1, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}},
		{{1, 0, 0}, {0, 0, 0}, {1, 1, 1}, {1, 1, 0}},
	}
	kValues := []int{2, 3, 3, 2, 1}

	for i := 0; i < len(matrixList); i++ {
		fmt.Printf("%d.\tInput matrix: \n\tmatrix = ", i+1)
		printMatrix(matrixList[i])
		fmt.Printf("\n\tk = %d\n", kValues[i])
		weakestRows := findKWeakestRows(matrixList[i], kValues[i])
		fmt.Printf("\n\tIndexes of the %d weakest rows: [%s]\n", kValues[i], formatIndexes(weakestRows))
		fmt.Println(strings.Repeat("-", 100))
	}
}
