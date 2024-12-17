package main

import (
	"fmt"
	"strings"
)

/*
	DisplayLinkedListWithForwardArrow method will display the linked list

not in the form of an array, but rather a list with arrows pointing to
the next element
*/
func DisplayLinkedListWithForwardArrow(l *EduLinkedListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(" → ")
		} else {
			fmt.Print(" → null")
		}
	}
}

type EduLinkedList struct {
	head *EduLinkedListNode
}

/*
InsertNodeAtHead method will insert a LinkedListNode at head

	of a linked list.
*/
func (l *EduLinkedList) InsertNodeAtHead(node *EduLinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

/*
	CreateLinkedList method will create the linked list using

the given integer array with the help of InsertAthead method.
*/
func (l *EduLinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(lst[i])
		l.InsertNodeAtHead(newNode)
	}
}

// DisplayLinkedList method will display the elements of linked list.
func (l *EduLinkedList) DisplayLinkedList() {
	temp := l.head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}

// Function to find the middle node of the linked list
func getMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {

	// Create two pointers, slow and fast ,initially pointing to the head
	slow := head
	fast := head

	// Traverse the linked list until fast reaches at the last node or NULL
	for fast != nil && fast.next != nil {

		// Move the slow pointer one step ahead
		slow = slow.next

		// Move the fast pointer two steps ahead
		fast = fast.next.next
	}

	// Return the slow pointer
	return slow
}

// Driver code
func main() {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{3, 2, 1},
		{10},
		{1, 2},
	}
	for i, input := range inputs {
		linkedList := new(EduLinkedList)
		linkedList.CreateLinkedList(input)
		fmt.Printf("%d.\tLinked list: ", i+1)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n\tMiddle of linked list is: %d\n", getMiddleNode(linkedList.head).data)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
