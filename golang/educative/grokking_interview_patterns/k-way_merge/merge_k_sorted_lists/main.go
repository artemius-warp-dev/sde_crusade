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
func DisplayLinkedListWithForwardArrow(head *EduLinkedListNode) {
	temp := head
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

func merge2Lists(head1, head2 *EduLinkedListNode) *EduLinkedListNode {
	dummy := &EduLinkedListNode{data: -1}

	prev := dummy

	for head1 != nil && head2 != nil {
		if head1.data <= head2.data {
			prev.next = head1
			head1 = head1.next
		} else {
			prev.next = head2
			head2 = head2.next
		}
		prev = prev.next
	}

	if head1 == nil {
		prev.next = head2
	} else {
		prev.next = head1
	}
	return dummy.next

}

func mergeKLists(lists []*EduLinkedList) *EduLinkedListNode {
	if len(lists) > 0 {
		step := 1
		for step < len(lists) {
			for i := 0; i < len(lists)-step; i += step * 2 {
				lists[i].head = merge2Lists(lists[i].head, lists[i+step].head)
			}
			step *= 2
		}
		return lists[0].head
	}
	return nil
}

// Driver code
func main() {
	inputLists := [][][]int{
		{
			{21, 23, 42},
			{1, 2, 4},
		},
		{
			{11, 41, 51},
			{21, 23, 42},
		},
		{
			{2},
			{1, 2, 4},
			{25, 56, 66, 72},
		},
		{
			{11, 41, 51},
			{2},
			{2},
			{2},
			{1, 2, 4},
		},
		{
			{10, 30},
			{15, 25},
			{1, 7},
			{3, 9},
			{100, 300},
			{115, 125},
			{10, 70},
			{30, 90},
		},
	}

	for index, value := range inputLists {
		fmt.Printf("%d.\tInput lists:\n", index+1)
		llLists := make([]*EduLinkedList, 0)
		for _, v := range value {
			a := new(EduLinkedList)
			a.CreateLinkedList(v)
			llLists = append(llLists, a)
			fmt.Printf("\t")
			DisplayLinkedListWithForwardArrow(a.head)
			fmt.Printf("\n")
		}
		fmt.Printf("\tMerges list:\n")
		fmt.Printf("\t")
		DisplayLinkedListWithForwardArrow(mergeKLists(llLists))
		fmt.Printf("\n")
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
