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

func reverse(head *EduLinkedListNode) *EduLinkedListNode {

	var prev, next *EduLinkedListNode
	curr := head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	head = prev
	return prev

}

func main() {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{3, 2, 1},
		{10},
		{1, 2},
	}

	for i := 0; i < len(input); i++ {
		inputLinkedList := EduLinkedList{}
		inputLinkedList.CreateLinkedList(input[i])
		fmt.Printf("%d. \tInput linked list: ", i+1)
		DisplayLinkedListWithForwardArrow(inputLinkedList.head)
		fmt.Printf("\n\tReversed linked list: ")
		DisplayLinkedListWithForwardArrow(reverse(inputLinkedList.head))
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
