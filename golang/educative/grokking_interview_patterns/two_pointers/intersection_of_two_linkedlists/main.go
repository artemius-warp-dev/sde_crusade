package main

import (
	"fmt"
	"strings"
)

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *ListNode
}

// NewLinkedList creates a linked list from a slice of integers
func NewLinkedList(values []int) *LinkedList {
	ll := &LinkedList{}
	if len(values) > 0 {
		ll.createLinkedList(values)
	}
	return ll
}

// createLinkedList initializes the linked list from a slice of values
func (ll *LinkedList) createLinkedList(values []int) {
	if len(values) == 0 {
		ll.Head = nil
		return
	}

	ll.Head = &ListNode{Val: values[0]}
	current := ll.Head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
}

// display prints the linked list
func display(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("None")
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a

}

// Driver code
func main() {
	listAValues := [][]int{
		{4, 1, 8, 4, 5},
		{1, 2, 3},
		{1, 2, 3},
		{2, 6, 4},
		{1},
		{2, 3, 6, 21, 5},
	}

	listBValues := [][]int{
		{5, 6, 6, 8, 4, 5},
		{4, 5, 6},
		{1, 2, 3},
		{1, 5, 4},
		{1},
		{7, 21, 5},
	}

	skipA := []int{2, 3, 0, 2, 0, 3}
	skipB := []int{3, 3, 0, 2, 0, 1}
	intersectVal := []int{8, 0, 1, 4, 1, 21}

	for i := 0; i < len(listAValues); i++ {
		listA := NewLinkedList(listAValues[i])
		listB := NewLinkedList(listBValues[i])

		// Create intersection if applicable
		if skipA[i] == 0 && skipB[i] == 0 {
			listB = listA
		} else if skipA[i] < len(listAValues[i]) && skipB[i] < len(listBValues[i]) {
			nodeA := listA.Head
			for j := 0; j < skipA[i]; j++ {
				nodeA = nodeA.Next
			}

			nodeB := listB.Head
			for j := 0; j < skipB[i]-1; j++ {
				nodeB = nodeB.Next
			}

			nodeB.Next = nodeA
		}

		result := getIntersectionNode(listA.Head, listB.Head)

		fmt.Printf("%d.\tList A: ", i+1)
		display(listA.Head)
		fmt.Printf("\tList B: ")
		display(listB.Head)
		fmt.Printf("\tskipA: %d\n", skipA[i])
		fmt.Printf("\tskipB: %d\n", skipB[i])
		fmt.Printf("\tintersect_val: %d\n", intersectVal[i])

		if result != nil {
			fmt.Printf("\n\tOutput: Intersected at %d\n", result.Val)
		} else {
			fmt.Printf("\n\tOutput: No intersection\n")
		}

		fmt.Println(strings.Repeat("-", 100))
	}
}
