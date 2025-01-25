package main

import (
	"fmt"
	"strings"
)

// LinkedListNode represents a node in the linked list
type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

// LinkedList represents the linked list
type EduLinkedList struct {
	head *EduLinkedListNode
}

// InsertNodeAtHead inserts a node at the head of the linked list
func (list *EduLinkedList) InsertNodeAtHead(data int) {
	newNode := &EduLinkedListNode{data: data, next: list.head}
	list.head = newNode
}

// CreateLinkedList creates a linked list from an array
func (list *EduLinkedList) CreateLinkedList(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		list.InsertNodeAtHead(arr[i])
	}
}

// PrintListWithForwardArrow prints the linked list with arrows
func (list *EduLinkedList) PrintListWithForwardArrow() {
	current := list.head
	for current != nil {
		fmt.Print(current.data)
		current = current.next
		if current != nil {
			fmt.Print(" → ")
		} else {
			fmt.Print(" → null")
		}
	}
	fmt.Println()
}

func DeleteNodes(head *EduLinkedListNode, m int, n int) *EduLinkedListNode {
	current := head
	var lastMNode *EduLinkedListNode

	for current != nil {
		mCount := m
		for current != nil && mCount > 0 {
			lastMNode = current
			current = current.next
			mCount--
		}

		nCount := n
		for current != nil && nCount > 0 {
			current = current.next
			nCount--
		}

		if lastMNode != nil {
			lastMNode.next = current
		}
	}

	return head
}

// Driver code
func main() {
	inputLists := [][]int{
		{5},
		{1, 2, 2, 3, 3, 3},
		{3, 7, 9},
		{10, 10, 100, 100, 100},
		{7, 7, 7, 7, 77, 77, 77, 77},
	}
	allMs := []int{1, 3, 1, 4, 5}
	allNs := []int{1, 1, 3, 2, 7}

	for i, arr := range inputLists {
		linkedList := EduLinkedList{}
		linkedList.CreateLinkedList(arr)

		fmt.Printf("%d.\tInput:\n\t", i+1)
		linkedList.PrintListWithForwardArrow()
		fmt.Printf("\tm = %d\n", allMs[i])
		fmt.Printf("\tn = %d\n", allNs[i])

		linkedList.head = DeleteNodes(linkedList.head, allMs[i], allNs[i])
		fmt.Println("\n\tOutput:")
		fmt.Print("\t")
		linkedList.PrintListWithForwardArrow()
		fmt.Println(strings.Repeat("-", 100))
	}
}
