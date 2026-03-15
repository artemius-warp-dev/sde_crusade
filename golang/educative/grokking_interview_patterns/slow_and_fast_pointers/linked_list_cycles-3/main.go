package main

import (
	"fmt"
	"strings"
)

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *ListNode
}

// NewLinkedList creates a linked list from a slice of integers
func NewLinkedList(values []int) *LinkedList {
	ll := &LinkedList{}
	if len(values) > 0 {
		ll.createLinkedList(values)
	}
	return ll
}

// createLinkedList initializes the linked list from a slice of values
func (ll *LinkedList) createLinkedList(values []int) {
	if len(values) == 0 {
		ll.Head = nil
		return
	}

	ll.Head = &ListNode{Val: values[0]}
	current := ll.Head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
}

// GetLength returns the number of nodes in the linked list
func (ll *LinkedList) GetLength(head *ListNode) int {
	temp := head
	length := 0
	for temp != nil {
		length++
		temp = temp.Next
	}
	return length
}

// GetNode returns the node at the specified position (index) of the linked list
func (ll *LinkedList) GetNode(head *ListNode, pos int) *ListNode {
	if pos != -1 {
		p := 0
		ptr := head
		for p < pos {
			ptr = ptr.Next
			p++
		}
		return ptr
	}
	return nil
}

func DisplayLinkedListWithForwardArrow(l *ListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.Val)
		temp = temp.Next
		if temp != nil {
			fmt.Print(" → ")
		} else {
			fmt.Print(" → null")
		}
	}
}
func DisplayLinkedListWithForwardArrowLoop(l *ListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.Val)
		temp = temp.Next
		if temp != nil {
			fmt.Print(" → ")
		}
	}
}

func countCycleLength(head *ListNode) int {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			length := 1
			slow = slow.Next

			for slow != fast {
				length++
				slow = slow.Next
			}
			return length
		}
	}
	return 0
}

// Driver code
func main() {
	inputs := [][]int{
		{2, 4, 6, 8, 10, 12},
		{1, 3, 5, 7, 9, 11},
		{0, 1, 2, 3, 4, 6},
		{3, 4, 7, 9, 11, 17},
		{5, 1, 4, 9, 2, 3},
	}
	pos := []int{0, -1, 1, -1, 2}
	j := 1

	for i := range inputs {
		inputLinkedList := &LinkedList{}
		inputLinkedList.createLinkedList(inputs[i])
		fmt.Printf("%d.\tInput: ", j)
		if pos[i] == -1 {
			DisplayLinkedListWithForwardArrow(inputLinkedList.Head)
		} else {
			DisplayLinkedListWithForwardArrowLoop(inputLinkedList.Head)
		}
		fmt.Println("\n\tpos:", pos[i])
		if pos[i] != -1 {
			length := inputLinkedList.GetLength(inputLinkedList.Head)
			lastNode := inputLinkedList.GetNode(inputLinkedList.Head, length-1)
			lastNode.Next = inputLinkedList.GetNode(inputLinkedList.Head, pos[i])
		}

		fmt.Printf("\n\tCycle length = %d\n", countCycleLength(inputLinkedList.Head))
		j++
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
