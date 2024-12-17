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
func DisplayLinkedListWithForwardArrowLoop(l *EduLinkedListNode) {
	temp := l
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(" → ")
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

// GetLength returns the number of nodes in the linked list
func (ll *EduLinkedList) GetLength(head *EduLinkedListNode) int {
	temp := head
	length := 0
	for temp != nil {
		length++
		temp = temp.next
	}
	return length
}

// GetNode returns the node at the specified position (index) of the linked list
func (ll *EduLinkedList) GetNode(head *EduLinkedListNode, pos int) *EduLinkedListNode {
	if pos != -1 {
		p := 0
		ptr := head
		for p < pos {
			ptr = ptr.next
			p++
		}
		return ptr
	}
	return nil
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

func detectCycle(head *EduLinkedListNode) bool {
	if head == nil {
		return false
	}

	// Initialize two pointers, slow and fast, to the head of the linked list
	slow, fast := head, head

	// Run the loop until we reach the end of the linked list or find a cycle
	for fast != nil && fast.next != nil {
		// Move the slow pointer one step at a time
		slow = slow.next
		// Move the fast pointer two steps at a time
		fast = fast.next.next

		// If there is a cycle, the slow and fast pointers will meet
		if slow == fast {
			return true
		}
	}

	// If we reach the end of the linked list and haven't found a cycle, return false
	return false
}

// Driver code
func main() {
	inputs := [][]int{
		{2, 4, 6, 8, 10, 12},
		{1, 3, 5, 7, 9, 11},
		{0, 1, 2, 3, 4, 6},
		{3, 4, 7, 9, 11, 17},
		{5, 1, 4, 9, 2, 3},
	}
	pos := []int{0, -1, 1, -1, 2}
	j := 1

	for i := range inputs {
		inputLinkedList := &EduLinkedList{}
		inputLinkedList.CreateLinkedList(inputs[i])
		fmt.Printf("%d.\tInput: ", j)
		if pos[i] == -1 {
			DisplayLinkedListWithForwardArrow(inputLinkedList.head)
		} else {
			DisplayLinkedListWithForwardArrowLoop(inputLinkedList.head)
		}
		fmt.Println("\n\tpos:", pos[i])
		if pos[i] != -1 {
			length := inputLinkedList.GetLength(inputLinkedList.head)
			lastNode := inputLinkedList.GetNode(inputLinkedList.head, length-1)
			lastNode.next = inputLinkedList.GetNode(inputLinkedList.head, pos[i])
		}

		fmt.Printf("\n\tDetected cycle = %t\n", detectCycle(inputLinkedList.head))
		j++
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
