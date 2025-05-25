package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

// Element holds the value, list index, and element index
type Element struct {
	value      int
	listIndex  int
	elementIdx int
}

// PriorityQueue implements heap.Interface and holds Elements
type PriorityQueue []Element

func (pq PriorityQueue) Len() int { return len(pq) }

// We want a min-heap based on Element.value
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Element))
}
l
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	elem := old[n-1]
	*pq = old[0 : n-1]
	return elem
}

func smallestRange(nums [][]int) []int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	maxVal := math.MinInt64

	rangeStart := 0
	rangeEnd := math.MaxInt64

	for i := 0; i < len(nums); i++ {
		heap.Push(pq, Element{value: nums[i][0], listIndex: i, elementIdx: 0})
		if nums[i][0] > maxVal {
			maxVal = nums[i][0]
		}
	}

	for pq.Len() == len(nums) {
		minElem := heap.Pop(pq).(Element)
		minVal := minElem.value
		row := minElem.listIndex
		col := minElem.elementIdx

		if maxVal-minVal < rangeEnd-rangeStart {
			rangeStart = minVal
			rangeEnd = maxVal
		}

		if col+1 < len(nums[row]) {
			nextVal := nums[row][col+1]
			heap.Push(pq, Element{value: nextVal, listIndex: row, elementIdx: col + 1})
			if nextVal > maxVal {
				maxVal = nextVal
			}
		} else {
			break
		}

	}

	return []int{rangeStart, rangeEnd}

}

func main() {
	testCases := [][][]int{
		{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}},
		{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
		{{1, 5}, {3, 7}, {4, 6}},
		{{1, 2}, {3, 4}, {8, 9}},
		{{1, 5, 10}, {2, 4, 6}, {3, 8, 15}},
	}

	for i, nums := range testCases {
		fmt.Printf("%d.\tnums: %v\n", i+1, nums)
		result := smallestRange(nums)
		fmt.Printf("\tSmallest Range: [%d, %d]\n", result[0], result[1])
		fmt.Println(strings.Repeat("-", 100))
	}
}
