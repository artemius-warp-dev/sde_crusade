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

func SwapTwoNodes(node1 *EduLinkedListNode, node2 *EduLinkedListNode) {
	temp := node1.data
	node1.data = node2.data
	node2.data = temp
}

func swapNodes(head *EduLinkedListNode, k int) *EduLinkedListNode {
	count := 0
	front := new(EduLinkedListNode)
	end := new(EduLinkedListNode)
	curr := head

	for curr != nil {
		count += 1
		if end != nil {
			end = end.next
		}
		if count == k {
			front = curr
			end = head
		}
		curr = curr.next
	}

	SwapTwoNodes(front, end)
	return head
}

// Driver code
func main() {
	inputList := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{6, 9, 3, 10, 7, 4, 6},
		{6, 9, 3, 4},
		{6, 2, 3, 6, 9},
		{6, 2},
	}

	k := []int{2, 3, 2, 3, 1}

	for index, value := range inputList {
		linkedList := new(EduLinkedList)
		linkedList.CreateLinkedList(value)
		fmt.Printf("%d.\tOriginal linked list is: ", index+1)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n\tk: %d\n", k[index])

		if k[index] <= 0 {
			fmt.Print("\tThe expected 'k' to have value from 1 to length of the linked list only.\n")
		} else {
			swapNodes(linkedList.head, k[index])
			fmt.Print("\tLinked list with swapped values: ")
			DisplayLinkedListWithForwardArrow(linkedList.head)
		}
		fmt.Printf("\n%s\n", strings.Repeat("-", 100))
	}
}
