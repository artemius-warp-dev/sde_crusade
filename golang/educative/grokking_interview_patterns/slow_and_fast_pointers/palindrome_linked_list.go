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

func ReverseList(slowPtr *EduLinkedListNode) *EduLinkedListNode {
	var prev, next, curr *EduLinkedListNode = nil, nil, slowPtr

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

func palindrome(head *EduLinkedListNode) bool {

	slow := head
	fast := head
	revertData := new(EduLinkedListNode)
	revertData = nil

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	revertData = ReverseList(slow)
	check := false
	check = compareTwoHalves(head, revertData)
	ReverseList(revertData)

	if check {
		return true
	}
	return false
}

func compareTwoHalves(firstHalf *EduLinkedListNode, secondHalf *EduLinkedListNode) bool {
	for firstHalf != nil && secondHalf != nil {
		if firstHalf.data != secondHalf.data {
			return false
		} else {
			firstHalf = firstHalf.next
			secondHalf = secondHalf.next
		}

	}
	return true
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

// Driver code
func main() {
	inputArray := [][]int{
		{2, 4, 6, 4, 2},
		{0, 3, 5, 5, 0},
		{9, 7, 4, 4, 7, 9},
		{5, 4, 7, 9, 4, 5},
		{5, 9, 8, 3, 8, 9, 5},
	}

	for i, input := range inputArray {
		linkedList := new(EduLinkedList)
		linkedList.CreateLinkedList(input)
		fmt.Printf("%d\tLinked List: ", i+1)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n")
		fmt.Printf("\tIs is Palindrome = ")
		if palindrome(linkedList.head) {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
