package main

import (
	"fmt"
	"strings"
)

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

func removeNthLastNode(head *EduLinkedListNode, n int) *EduLinkedListNode {

	left := head
	right := left

	for i := 0; i < n; i++ {
		right = right.next
	}

	if right == nil {

		head = head.next
	} else {
		for right.next != nil {
			right = right.next
			left = left.next
		}
		left.next = left.next.next
	}

	return head
}

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

// Driver Code
func main() {
	inputs := [][]int{
		{23, 89, 10, 5, 67, 39, 70, 28},
		{34, 53, 6, 95, 38, 28, 17, 63, 16, 76},
		{288, 224, 275, 390, 4, 383, 330, 60, 193},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{69, 8, 49, 106, 116, 112, 104, 129, 39, 14, 27, 12},
	}

	n := []int{4, 1, 6, 9, 11}

	for i := 0; i < len(inputs); i++ {
		inputLinkedList := EduLinkedList{}
		inputLinkedList.CreateLinkedList(inputs[i])
		fmt.Printf("%d.\tLinked List:\t\t", i+1)
		DisplayLinkedListWithForwardArrow(inputLinkedList.head)
		fmt.Printf("\n\tn = %d", n[i])
		fmt.Printf("\n\tUpdated Linked List:\t")
		DisplayLinkedListWithForwardArrow(removeNthLastNode(inputLinkedList.head, n[i]))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
