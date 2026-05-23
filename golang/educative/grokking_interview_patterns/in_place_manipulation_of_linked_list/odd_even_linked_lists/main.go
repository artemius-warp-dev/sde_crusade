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

// display prints the linked list
func display(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("None")
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

func main() {
	inputLists := [][]int{
		{1, 1, 2, 2, 3, -1, 10, 12},
		{10, 20, -22, 21, -12},
		{1, 1, 1},
		{-2, -5, -6, 0, -1, -4},
		{3, 1, 5, 7, -4, -2, -1, -6},
	}

	for i, inputList := range inputLists {
		obj := new(LinkedList)
		obj.createLinkedList(inputList)
		fmt.Printf("%d.\tOriginal list: ", i+1)
		display(obj.Head)
		oddEvenList(obj.Head)
		fmt.Print("\n\tAfter folding: ")
		display(obj.Head)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
