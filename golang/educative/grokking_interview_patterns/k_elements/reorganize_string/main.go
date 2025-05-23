package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Set struct {
	count int
	char  rune
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
	return h[i].count > h[j].count
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

func reorganizeString(str string) string {
	charCounter := make(map[rune]int)
	for _, c := range str {
		freq := 1
		if _, ok := charCounter[c]; ok {
			freq = charCounter[c] + 1
		}
		charCounter[c] = freq
	}

	mostFreqChars := newMaxHeap()

	for i, _ := range charCounter {
		heap.Push(mostFreqChars, Set{char: i, count: charCounter[i]})
	}

	result := ""
	previous := new(Set)

	for !mostFreqChars.Empty() || previous.count != 0 && previous.char != 0 {
		if (previous.count != 0 && previous.char != 0) && mostFreqChars.Empty() {
			return ""
		}

		poppedElement := heap.Pop(mostFreqChars).(Set)
		count, char := poppedElement.count, poppedElement.char
		result += string(char)

		count -= 1

		if previous.count != 0 && previous.char != 0 {
			heap.Push(mostFreqChars, *previous)
			previous = new(Set)
		}

		if count != 0 {
			previous.count = count
			previous.char = char
		}
	}

	return result
}

// Driver code
func main() {
	testCases := []string{"programming", "hello", "fofjjb", "abbacdde", "aba", "awesome", "aaab"}

	for i, s := range testCases {
		fmt.Printf("%d.\tInput string: \"%s\"\n", i+1, s)
		fmt.Printf("\tReorganized string: \"%s\"\n", reorganizeString(s))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
