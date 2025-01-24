package main

import (
	"fmt"
)

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

func reverseKGroups(head *EduLinkedListNode, k int) *EduLinkedListNode {

	dummy := &EduLinkedListNode{data: 0}
	dummy.next = head
	ptr := dummy

	for ptr != nil {
		tracker := ptr
		for i := 0; i < k; i++ {
			if tracker == nil {
				break
			}

			tracker = tracker.next

		}

		if tracker == nil {
			break
		}

		previous, current := reverseLinkedList(ptr.next, k)

		lastNodeOfReversedGroup := ptr.next
		lastNodeOfReversedGroup.next = current
		ptr.next = previous
		ptr = lastNodeOfReversedGroup
	}

	return dummy.next

}

func reverseLinkedList(head *EduLinkedListNode, k int) (*EduLinkedListNode, *EduLinkedListNode) {

	var previous, current, next *EduLinkedListNode = nil, head, nil
	for i := 0; i < k; i++ {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	return previous, current

}

// Driver code
func main() {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	k := 3
	var inputLinkedList EduLinkedList
	inputLinkedList.CreateLinkedList(inputList)

	fmt.Print("Linked list: ")
	DisplayLinkedListWithForwardArrow(inputLinkedList.head)
	fmt.Print("\n\n")

	fmt.Print("Reversed list: ")
	DisplayLinkedListWithForwardArrow(reverseKGroups(inputLinkedList.head, k))
	fmt.Print("\n\n")
}
