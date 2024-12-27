package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Deque struct {
	items *list.List
}

// NewDeque is a constructor that will declare and return the Deque type object
func NewDeque() *Deque {
	return &Deque{list.New()}
}

// PushFront will push an element at the front of the dequeue
func (d *Deque) PushFront(val int) {
	d.items.PushFront(val)
}

// PushBack will push an element at the back of the dequeue
func (d *Deque) PushBack(val int) {
	d.items.PushBack(val)
}

// PopFront will pop an element from the front of the dequeue
func (d *Deque) PopFront() int {
	return d.items.Remove(d.items.Front()).(int)
}

// PopBack will pop an element from the back of the dequeue
func (d *Deque) PopBack() int {
	return d.items.Remove(d.items.Back()).(int)
}

// Front will return the element from the front of the dequeue
func (d *Deque) Front() int {
	return d.items.Front().Value.(int)
}

// Back will return the element from the back of the dequeue
func (d *Deque) Back() int {
	return d.items.Back().Value.(int)
}

// Empty will check if the dequeue is empty or not
func (d *Deque) Empty() bool {
	return d.items.Len() == 0
}

// Len will return the length of the dequeue
func (d *Deque) Len() int {
	return d.items.Len()
}

func (d *Deque) Print() string {
	temp := d.items.Front()
	s := "["
	for temp != nil {
		temp2, _ := temp.Value.(int)
		s += strconv.Itoa(temp2)
		temp = temp.Next()
		if temp != nil {
			s += " , "
		}
	}
	s += "]"
	return s
}

func cleanUp(i int, currentWindow *Deque, nums []int) {
	for currentWindow.Len() > 0 && nums[i] >= nums[currentWindow.Back()] {
		currentWindow.PopBack()
	}
}

func findMaxSlidingWindow(nums []int, w int) []int {
	n := len(nums)
	if n == 1 {
		return nums
	}

	output := make([]int, n-w+1)
	currentWindow := NewDeque()

	for i := 0; i < w; i++ {
		cleanUp(i, currentWindow, nums)
		currentWindow.PushBack(i)
	}

	output[0] = nums[currentWindow.Front()]

	for i := w; i < n; i++ {
		cleanUp(i, currentWindow, nums)

		if currentWindow.Len() > 0 && currentWindow.Front() <= i-w {
			currentWindow.PopFront()
		}

		currentWindow.PushBack(i)

		output[i-w+1] = nums[currentWindow.Front()]

	}
	return output
}

// driver code
func main() {
	windowSizes := []int{3, 3, 3, 3, 2, 4, 3, 2, 3, 6}
	numsList := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{1, 5, 8, 10, 10, 10, 12, 14, 15, 19, 19, 19, 17, 14, 13, 12, 12, 12, 14, 18, 22, 26, 26, 26, 28, 29, 30},
		{10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67},
		{4, 5, 6, 1, 2, 3},
		{9, 5, 3, 1, 6, 3},
		{2, 4, 6, 8, 10, 12, 14, 16},
		{-1, -1, -2, -4, -6, -7},
		{4, 4, 4, 4, 4, 4},
	}

	for index, value := range numsList {
		fmt.Printf("%d.\tInput array:\t%s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\tWindow size:\t%d\n", windowSizes[index])
		output := findMaxSlidingWindow(value, windowSizes[index])
		fmt.Printf("\n\tMaximum in each sliding window:\t%s\n", strings.Replace(fmt.Sprint(output), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
