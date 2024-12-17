package main

import (
	"fmt"
	"strings"
)

// Definition of a binary tree node
// type LinkedListNode[T any] struct {
//     Data T
//     Next *LinkedListNode[T]
// }

// func twinSum(head *LinkedListNode[int]) int {

// 	// Replace this placeholder return statement with your code
// 	slow, fast := head, head
// 	for fast.Next != nil {
// 		slow = slow.Next
// 		if fast.Next.Next == nil {
// 			fast = fast.Next
// 		} else {
// 			fast = fast.Next.Next
// 		}

// 	}

// 	var prev, next, curr *LinkedListNode[int] = nil, nil, slow

// 	for curr != nil {
// 		next = curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	curr = head

// 	maxSum := 0
// 	sum := 0
// 	for prev != nil {
// 		sum = curr.Data + prev.Data
// 		if sum > maxSum {
// 			maxSum = sum
// 		}
// 		curr = curr.Next
// 		prev = prev.Next
// 	}

// 	return maxSum
// }

// Define a generic LinkedList[T] structure
type LinkedList[T any] struct {
	Head *LinkedListNode[T]
}

/*
InsertNodeAtHead method will insert a LinkedListNode at head

	of a linked list.
*/
func (list *LinkedList[T]) InsertNodeAtHead(node *LinkedListNode[T]) {
	if list.Head == nil {
		list.Head = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
}

/*
	CreateLinkedList method will create the linked list using

the given integer array with the help of InsertAthead method.
*/
func (list *LinkedList[T]) CreateLinkedList(lst []T) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := &LinkedListNode[T]{Data: lst[i], Next: nil}
		list.InsertNodeAtHead(newNode)
	}
}

// LinkedListNode represents a node in the linked list
type LinkedListNode[T any] struct {
	Data T
	Next *LinkedListNode[T]
}

/*
	DisplayLinkedListWithForwardArrow method will display the linked list

not in the form of an array, but rather a list with arrows pointing to
the next element
*/
func PrintListWithForwardArrow(node *LinkedListNode[int]) {
	temp := node
	for temp != nil {
		fmt.Print(temp.Data)
		temp = temp.Next
		if temp != nil {
			fmt.Print(" → ")
		} else {
			fmt.Print(" → null")
		}
	}
}

func twinSum(head *LinkedListNode[int]) int {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr, prev := slow, (*LinkedListNode[int])(nil)

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	maxSum := 0
	curr = head

	for prev != nil {
		sum := curr.Data + prev.Data
		if sum > maxSum {
			maxSum = sum
		}
		prev = prev.Next
		curr = curr.Next
	}

	return maxSum
}

// Driver code
func main() {
	lists := [][]int{
		{2, 3, 5, 7},
		{81, 144, 64, 121, 25, 49},
		{4, 5, 6, 7},
		{1, 1000},
		{11, 77, 44, 99, 22, 66, 55, 88},
	}

	for i, list := range lists {
		inputLinkedList := new(LinkedList[int])
		inputLinkedList.CreateLinkedList(list)
		fmt.Printf("%d.\tLinked list: ", i+1)
		PrintListWithForwardArrow(inputLinkedList.Head)
		fmt.Printf("\n\tMaximum twin sum: %d", twinSum(inputLinkedList.Head))
		fmt.Println("\n" + strings.Repeat("-", 100))
	}
}
