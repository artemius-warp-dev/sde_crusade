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

func splitListToParts(head *EduLinkedListNode, k int) []*EduLinkedListNode {
	ans := make([]*EduLinkedListNode, k)

	size := 0
	current := head
	for current != nil {
		size++
		current = current.next
	}

	split := size / k
	remaining := size % k

	current = head
	var prev *EduLinkedListNode

	for i := 0; i < k; i++ {
		ans[i] = current
		currentSize := split
		if remaining > 0 {
			currentSize++
			remaining--
		}

		for j := 0; j < currentSize; j++ {
			prev = current
			if current != nil {
				current = current.next
			}
		}

		if prev != nil {
			prev.next = nil
		}

	}
	printLinkedListNodes(ans[0])
	return ans

}

func printLinkedListNodes(node *EduLinkedListNode) {
	if node == nil {
		fmt.Println()
		return
	}
	fmt.Printf("%v ", node.data)
	printLinkedListNodes(node.next)
}

// Main function to test the splitListToParts function
func main() {
	lists := [][]int{
		{9, 7, 8, 7, 7, 6},
		{2, 3, 5, 7, 11},
		{4, 4, 4, 4, 4},
		{1, 2, 3, 11, 22, 33},
		{1, 2, 6, 3, 4, 5, 6},
	}

	ks := []int{7, 2, 4, 3, 6}

	for i, dataues := range lists {
		inputLinkedList := &EduLinkedList{}
		inputLinkedList.CreateLinkedList(dataues)
		fmt.Printf("%d. Linked list: ", i+1)
		DisplayLinkedListWithForwardArrow(inputLinkedList.head)

		fmt.Printf("\n   k: %d\n", ks[i])
		result := splitListToParts(inputLinkedList.head, ks[i])

		fmt.Println("   Linked list parts:")
		for j, part := range result {
			if part != nil {
				fmt.Print("   [")
				current := part
				for current != nil {
					fmt.Print(current.data)
					if current.next != nil {
						fmt.Print(", ")
					}
					current = current.next
				}
				fmt.Print("]")
			} else {
				fmt.Print("[]")
			}
			if j < len(result)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
		fmt.Println(strings.Repeat("-", 100))

	}
}
