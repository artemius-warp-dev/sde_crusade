package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func kSmallestNumber(lists [][]int, k int) int {

	minElements := newMinHeap()

	for index, list := range lists {
		if len(list) == 0 {
			continue
		} else {
			heap.Push(minElements, Set{element: lists[index][0], list_index: index, element_index: 0})
		}
	}

	numberCount, smallestNumber := 0, 0

	for minElements.Empty() == false {
		popValue := heap.Pop(minElements).(Set)
		smallestNumber = popValue.element
		listIndex, elemIndex := popValue.list_index, popValue.element_index
		numberCount++

		if numberCount == k {
			break
		}

		if elemIndex+1 < len(lists[listIndex]) {
			heap.Push(minElements, Set{element: lists[listIndex][elemIndex+1], list_index: listIndex, element_index: elemIndex + 1})
		}

	}

	return smallestNumber
}

// Driver code
func main() {
	lists := [][][]int{
		{{2, 6, 8}, {3, 6, 10}, {5, 8, 11}},
		{{1, 2, 3}, {4, 5}, {6, 7, 8, 15}, {10, 11, 12, 13}, {5, 10}},
		{{}, {}, {}},
		{{1, 1, 3, 8}, {5, 5, 7, 9}, {3, 5, 8, 12}},
		{{5, 8, 9, 17}, {}, {8, 17, 23, 24}},
	}
	k := []int{5, 50, 7, 4, 8}

	for i, list := range lists {
		fmt.Printf("%d.\t Input lists: %s \n\t K = %d\n", i+1, strings.Replace(fmt.Sprint(list), " ", ", ", -1), k[i])
		fmt.Printf("\t %dth smallest number from the given lists is: %d\n", k[i], kSmallestNumber(list, k[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
