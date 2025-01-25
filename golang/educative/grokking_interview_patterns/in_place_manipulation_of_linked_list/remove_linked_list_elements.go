package main

import (
	"fmt"
	"strings"
)

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
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

func RemoveElements(head *EduLinkedListNode, k int) *EduLinkedListNode {
	dummy := &EduLinkedListNode{data: 0, next: head}

	// for head != nil && head.data == k {
	// 	head = head.next
	// }

	prev := dummy
	curr := head

	for curr != nil {
		if curr.data == k {
			prev.next = curr.next
			curr = curr.next
		} else {
			prev = curr
			curr = curr.next
		}
	}
	return dummy.next
}

func main() {
	// Test cases
	lists := [][]int{
		{9, 7, 8, 7, 7, 6},
		{2, 3, 5, 7, 11},
		{4, 4, 4, 4, 4},
		{1, 2, 3, 11, 22, 33},
		{1, 2, 6, 3, 4, 5, 6},
	}

	ks := []int{7, 8, 4, 3, 6}

	for i, values := range lists {
		inputLinkedList := &EduLinkedList{}
		inputLinkedList.CreateLinkedList(values)
		fmt.Printf("%d.\tLinked list: ", i+1)
		DisplayLinkedListWithForwardArrow(inputLinkedList.head)

		fmt.Printf("\n\tk: %d\n", ks[i])

		fmt.Printf("\tLinked list after removing elements: ")
		res := RemoveElements(inputLinkedList.head, ks[i])
		DisplayLinkedListWithForwardArrow(res)
		fmt.Println("\n")
		fmt.Println(strings.Repeat("-", 100))
	}
}
