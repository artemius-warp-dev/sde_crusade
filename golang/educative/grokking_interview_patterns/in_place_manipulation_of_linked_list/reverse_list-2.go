package main

import (
	"fmt"
	"strings"
)

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

// func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
// 	node := new(EduLinkedListNode)
// 	node.data = data
// 	node.next = next
// 	return node
// }

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

func reverseBetween(head *EduLinkedListNode, left int, right int) *EduLinkedListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &EduLinkedListNode{data: 0}
	dummy.next = head
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.next
	}

	curr := prev.next

	for i := 0; i < right-left; i++ {
		nextNode := curr.next
		curr.next = nextNode.next
		nextNode.next = prev.next
		prev.next = nextNode
	}
	return dummy.next

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
	left := []int{1, 3, 2, 1, 1}
	right := []int{5, 6, 4, 3, 2}

	for index, value := range inputList {
		inputLinkedList := new(EduLinkedList)
		inputLinkedList.CreateLinkedList(value)
		fmt.Printf("%d.\tOriginal linked list: ", index+1)
		DisplayLinkedListWithForwardArrow(inputLinkedList.head)
		fmt.Println("\n\tleft: ", left[index], "right: ", right[index])
		if left[index] <= 0 {
			fmt.Printf("\tThe expected 'left' and 'right' to have value from 1 to length of the linked list only\n")
		} else {
			result := reverseBetween(inputLinkedList.head, left[index], right[index])
			fmt.Printf("\n\tReversed linked list: ")
			DisplayLinkedListWithForwardArrow(result)
			fmt.Printf("\n")
			fmt.Printf("%s\n", strings.Repeat("-", 100))
		}
	}
}
