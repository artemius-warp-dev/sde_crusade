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

func reorderList(head *EduLinkedListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow, fast = slow.next, fast.next.next
	}

	prev := new(EduLinkedListNode)
	prev = nil
	curr := slow
	tmp := new(EduLinkedListNode)
	tmp = nil

	for curr != nil {
		tmp = curr.next
		curr.next = prev
		prev = curr
		curr = tmp
	}

	first, second := head, prev

	for second.next != nil {
		tmp = first.next
		first.next = second
		first = tmp

		tmp = second.next
		second.next = first
		second = tmp
	}

}

// Driver code
func main() {
	inputLists := [][]int{
		{1, 1, 2, 2, 3, -1, 10, 12},
		{10, 20, -22, 21, -12},
		{1, 1, 1},
		{-2, -5, -6, 0, -1, -4},
		{3, 1, 5, 7, -4, -2, -1, -6},
	}

	for i, inputList := range inputLists {
		obj := new(EduLinkedList)
		obj.CreateLinkedList(inputList)
		fmt.Printf("%d.\tOriginal list: ", i+1)
		DisplayLinkedListWithForwardArrow(obj.head)
		reorderList(obj.head)
		fmt.Print("\n\tAfter folding: ")
		DisplayLinkedListWithForwardArrow(obj.head)
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
