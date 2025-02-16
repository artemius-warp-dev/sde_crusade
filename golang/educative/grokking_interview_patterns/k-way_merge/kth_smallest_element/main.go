package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

// Template for the min heap

type Set struct {
	elem int
	i    int
	j    int
}

type MinHeap []Set

// newMinHeap initializes an instance of MinHeap
func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

// Len function returns the length of MinHeap
func (h MinHeap) Len() int {
	return len(h)
}

// Less function compares the two elements of MinHeap given their indices
func (h MinHeap) Less(i, j int) bool {
	return h[i].elem < h[j].elem
}

// Swap function swaps the values of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push function pushes the given element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

// Pop function pops the top element of MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallestElement(matrix [][]int, k int) int {
	rowCount := len(matrix)

	minNumbers := &MinHeap{}

	for index := 0; index < int(math.Min(float64(rowCount), float64(k))); index++ {
		heap.Push(minNumbers, Set{elem: matrix[index][0], i: index, j: 0})
	}

	numbersChecked := 0
	var smallestElement int
	for minNumbers.Len() > 0 {
		s := heap.Pop(minNumbers).(Set)
		smallestElement = s.elem
		var rowIndex = s.i
		var colIndex = s.j
		numbersChecked++
		if numbersChecked == k {
			break
		}

		if colIndex+1 < len(matrix[rowIndex]) {
			heap.Push(minNumbers, Set{elem: matrix[rowIndex][colIndex+1], i: rowIndex, j: colIndex + 1})
		}
	}
	return smallestElement
}

// Driver code
func main() {
	matrix := [][][]int{
		{{2, 6, 8},
			{3, 7, 10},
			{5, 8, 11}},

		{{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}},

		{{5}},

		{{2, 5, 7, 9, 10},
			{4, 6, 8, 12, 14},
			{11, 13, 16, 18, 20},
			{15, 17, 21, 24, 26},
			{19, 22, 23, 25, 28}},

		{{3, 5, 7, 9, 11, 13},
			{6, 8, 10, 12, 14, 16},
			{15, 17, 19, 21, 23, 25},
			{18, 20, 22, 24, 26, 28},
			{27, 29, 31, 33, 35, 37},
			{30, 32, 34, 36, 38, 40}},
	}
	k := []int{3, 4, 1, 10, 15}

	for i, tempMatrix := range matrix {
		fmt.Printf("%d.\t Input matrix: %s \n\t K = %d\n", i+1, strings.Replace(fmt.Sprint(tempMatrix), " ", ", ", -1), k[i])
		fmt.Printf("\t %dth smallest number from the given matrix is: %d\n", k[i], kthSmallestElement(tempMatrix, k[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
