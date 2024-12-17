package main

import (
	"fmt"
	"strings"
)

// func SplitCircularLinkedList(head *EduLinkedListNode) []*EduLinkedListNode {
// 	//size := 0
// 	slow, fast := head, head.next.next
// 	for fast != head {
// 		slow = slow.next
// 		fast = fast.next
// 		if fast != head {
// 			fast = fast.next
// 		} else {
// 			break
// 		}

// 	}

// 	head1, head2 := head, slow.next
// 	slow.next = head1
// 	ptr := head2
// 	for ptr.next != head {
// 		ptr = ptr.next
// 	}
// 	ptr.next = head2
// 	result := []*EduLinkedListNode{head1, head2}
// 	fmt.Printf("%+v\n", head1)
// 	fmt.Printf("%+v\n", head2)

// 	return result
// }

// EduLinkedList represents a circular linked list
type EduLinkedList struct {
	head *EduLinkedListNode
}

// NewEduLinkedList creates a new circular linked list
func NewEduLinkedList() *EduLinkedList {
	return &EduLinkedList{}
}

// InsertNodeAtHead inserts a new node at the head of the circular linked list
func (l *EduLinkedList) InsertNodeAtHead(node *EduLinkedListNode) {
	if l.head == nil {
		// If the list is empty, the new node points to itself
		l.head = node
		node.next = node
	} else {
		// Insert at head and make the last node point to the new head
		last := l.head
		for last.next != l.head {
			last = last.next
		}
		node.next = l.head
		l.head = node
		last.next = l.head // Update last node to point to new head
	}
}

// CreateLinkedList creates the circular linked list using an integer slice
func (l *EduLinkedList) CreateLinkedList(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(arr[i])
		l.InsertNodeAtHead(newNode)
	}
}

// ToArray converts the circular linked list to a slice
func (l *EduLinkedList) ToArray() []int {
	if l.head == nil {
		return []int{}
	}

	seenNodes := make(map[*EduLinkedListNode]bool)
	result := []int{}
	current := l.head

	for !seenNodes[current] {
		result = append(result, current.data)
		seenNodes[current] = true
		current = current.next
	}

	return result
}

// PrintList provides methods to print the circular linked list
type PrintList struct{}

// PrintCircularLinkedList prints the circular linked list
func (p *PrintList) PrintCircularLinkedList(head *EduLinkedListNode) {
	if head == nil {
		fmt.Println("List is empty")
		return
	}

	seenNodes := make(map[*EduLinkedListNode]bool)
	temp := head

	for temp != nil && !seenNodes[temp] {
		fmt.Printf("%d → ", temp.data)
		seenNodes[temp] = true
		temp = temp.next
		if temp == head { // When we come back to the head node
			fmt.Print("(head) ")
			break
		}
	}
	fmt.Println() // Move to the next line after printing
}

// LinkedListToArray converts the linked list to a slice
func (p *PrintList) LinkedListToArray(head *EduLinkedListNode) []int {
	if head == nil {
		return []int{}
	}

	seenNodes := make(map[*EduLinkedListNode]bool)
	result := []int{}
	current := head

	for !seenNodes[current] {
		result = append(result, current.data)
		seenNodes[current] = true
		current = current.next
	}

	return result
}

// EduLinkedListNode represents a node in the circular linked list
type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

// NewLinkedListNode creates a new EduLinkedListNode
func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = next
	return node
}

// InitLinkedListNode initializes a new EduLinkedListNode with no next node
func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}

// SplitCircularLinkedList splits the circular linked list into two halves
func SplitCircularLinkedList(head *EduLinkedListNode) []*EduLinkedListNode {
	// Initialize the slow and fast pointer at the head
	slow := head
	fast := head

	// Use slow and fast pointers to find the midpoint
	// Move the slow pointer by one step and the fast pointer by two steps in each iteration.
	// This loop stops when fast reaches or almost reaches the end of the list,
	// indicating that slow is at the midpoint.
	for fast.next != head && fast.next.next != head {
		slow = slow.next
		fast = fast.next.next
	}

	// Set head1 to the start of the first half and head2 to the start of the second half
	head1 := head
	head2 := slow.next
	slow.next = head1 // Complete the first half
	// Adjust the fast pointer to complete the second half's circular structure
	fast = head2

	// Complete the second half
	for fast.next != head {
		fast = fast.next
	}
	fast.next = head2 // Point last node of second half to head2

	return []*EduLinkedListNode{head1, head2}
}

// Driver code
func main() {
	lists := [][]int{
		{1, 5, 7},
		{2, 6, 1, 5},
		{3, 1, 4, 2, 5},
		{8, 10, 12, 14, 16, 18},
		{9, 10},
	}

	printList := &PrintList{}

	for index, list := range lists {
		inputLinkedList := NewEduLinkedList()
		inputLinkedList.CreateLinkedList(list)

		fmt.Printf("%d. Linked list: ", index+1)
		printList.PrintCircularLinkedList(inputLinkedList.head)

		// Get the split lists
		splitLists := SplitCircularLinkedList(inputLinkedList.head) // Capture as a slice
		splitList1 := splitLists[0]
		splitList2 := splitLists[1]

		// Convert each split list to arrays for printing
		arrSplitList1 := printList.LinkedListToArray(splitList1)
		arrSplitList2 := printList.LinkedListToArray(splitList2)

		// Print the split lists in the desired format
		fmt.Printf("   Split Lists: [[%v], [%v]]\n", arrSplitList1, arrSplitList2)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
