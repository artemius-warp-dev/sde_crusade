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
			fmt.Print(" → null\n")
		}
	}
}

func reverseEvenLengthGroups(head *EduLinkedListNode) {
	prev := head
	groupLen := 2
	for prev.next != nil {
		node := prev
		numNodes := 0
		for i := 0; i < groupLen; i++ {
			if node.next == nil {
				break
			}
			numNodes += 1
			node = node.next
		}

		if numNodes%2 == 0 {
			reverse := node.next
			curr := prev.next
			for i := 0; i < numNodes; i++ {
				currNext := curr.next
				curr.next = reverse
				reverse = curr
				curr = currNext
			}
			prevNext := prev.next
			prev.next = node
			prev = prevNext
		}
		prev = node
	}
	groupLen += 1

}

func main() {
	inputList := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{15},
		{16, 17},
	}

	for index, value := range inputList {
		linkedList := new(EduLinkedList)
		linkedList.CreateLinkedList(value)
		fmt.Printf("%d.\tIf we reverse the even length groups of the linked list\n\t\t", index+1)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Printf("\n\tWe will get:\n\t\t")
		reverseEvenLengthGroups(linkedList.head)
		DisplayLinkedListWithForwardArrow(linkedList.head)
		fmt.Print(strings.Repeat("-", 100), "\n")
	}
}
