package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Structure for MaxHeap
type MaxHeap []int

// newMaxHeap function intializes an instance of the MaxHeap
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

// Greater returns true if the first of the given elements is greater than the second one
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap function swaps the values at the given indices
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() int {
	return h[0]
}

// Push function inserts the element in the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MinHeap structure initialization
type MinHeap []int

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

// Less function compares two elements of the MinHeap given their indexes
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap function swaps the values of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() int {
	return h[0]
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianOfStream struct {
	maximumList MaxHeap
	minimumList MinHeap
}

func (this *MedianOfStream) Init() {

	this.maximumList = MaxHeap{}
	this.minimumList = MinHeap{}

}

func (this *MedianOfStream) InsertNum(num int) float64 {
	if this.maximumList.Empty() || this.maximumList.Top() >= num {
		heap.Push(&this.maximumList, num)
	} else {
		heap.Push(&this.minimumList, num)
	}

	if this.maximumList.Len() > this.minimumList.Len()+1 {
		heap.Push(&this.minimumList, heap.Pop(&this.maximumList).(int))
	} else if this.maximumList.Len() < this.minimumList.Len() {
		heap.Push(&this.maximumList, heap.Pop(&this.minimumList).(int))
	}

	return this.FindMedian()
}

func (this *MedianOfStream) FindMedian() float64 {
	if this.maximumList.Len() == this.minimumList.Len() {
		return float64(this.maximumList.Top())/2.0 + float64(this.minimumList.Top())/2.0

		//return (float64(this.maximumList.Top()) + float64(this.minimumList.Top())) / 2.0
	}

	return float64(this.maximumList.Top())
}

func main() {
	nums := []int{35, 22, 30, 25, 1}
	numList := make([]int, 0)
	medianOfStream := MedianOfStream{}
	for index, value := range nums {
		numList = append(numList, value)
		fmt.Printf("%d.\tData stream: %s\n", index+1, strings.Replace(fmt.Sprint(numList), " ", ", ", -1))
		fmt.Printf("\tThe median for the given numbers is: %.1f\n", medianOfStream.InsertNum(value))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
