package main

import (
	"fmt"
	"strings"
)

// EduLinkedListNode represents a node in the linked list
type LinkedListNode[T any] struct {
	Data T
	Next *LinkedListNode[T]
}

// Define a generic LinkedList[T] structure
type LinkedList[T any] struct {
	Head *LinkedListNode[T]
}

// Method to insert a LinkedListNode at the head of the linked list
func (list *LinkedList[T]) InsertNodeAtHead(node *LinkedListNode[T]) {
	if list.Head != nil {
		node.Next = list.Head
	}
	list.Head = node
}

// Method to create the linked list using the given slice
func (list *LinkedList[T]) CreateLinkedList(arr []T) {
	for i := len(arr) - 1; i >= 0; i-- {
		newNode := &LinkedListNode[T]{Data: arr[i], Next: nil}
		list.InsertNodeAtHead(newNode)
	}
}

// Method to print the elements of the linked list
func (list *LinkedList[T]) PrintListWithForwardArrow() {
	temp := list.Head
	for temp != nil {
		fmt.Print(temp.Data)
		temp = temp.Next
		if temp != nil {
			fmt.Print(" → ")
		} else {
			fmt.Print(" → null")
		}
	}
	fmt.Println()
}

func RemoveDuplicates(Head *LinkedListNode[int]) *LinkedListNode[int] {
	current := Head
	for current != nil && current.Next != nil {
		if current.Next.Data == current.Data {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return Head
}

// Driver code
func main() {
	inputLists := [][]int{
		{1, 2, 2, 3, 3, 3},
		{-21, -21, -21, -21, -21, -21, -21},
		{3, 7, 9},
		{-100, -100, -100, -10, -10, 0, 10, 10, 100, 100, 100},
		{-77, -77, -7, -7, -7, -7, 7, 7, 7, 7, 77, 77, 77, 77},
	}

	for i, inputList := range inputLists {
		ll := &LinkedList[int]{}
		ll.CreateLinkedList(inputList)

		fmt.Printf("%d.\tInput: ", i+1)
		ll.PrintListWithForwardArrow()

		fmt.Printf("\n\tOutput: ")
		ll.Head = RemoveDuplicates(ll.Head)
		ll.PrintListWithForwardArrow()

		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
