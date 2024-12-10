package main

import "fmt"

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

	slow, fast := head, head

	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	reversedHalfPtr := ReverseList(slow)
	slow = head
	for reversedHalfPtr.next != nil {
		if slow.data != reversedHalfPtr.data {
			return false
		}

		slow = slow.next
		reversedHalfPtr = reversedHalfPtr.next
	}

	return true
}

func main() {
	data := []int{1, 2, 4, 4, 2, 1}
	var list EduLinkedList
	list.CreateLinkedList(data)
	list.DisplayLinkedList()
	fmt.Println("\n", palindrome(list.head))
}
