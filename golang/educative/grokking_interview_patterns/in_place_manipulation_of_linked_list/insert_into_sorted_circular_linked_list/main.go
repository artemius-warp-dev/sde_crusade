package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

// Constructor for Node (directly creating an instance of Node)
func NewNode(val int, next *Node) *Node {
	return &Node{Val: val, Next: next}
}

// EduLinkedList struct to define the linked list
type EduLinkedList struct {
	head *Node
}

// Constructor to initialize with values
func NewEduLinkedList(values []int) *EduLinkedList {
	eduList := &EduLinkedList{head: nil}
	eduList.createLinkedList(values)
	return eduList
}

// Function to create a linked list from a list of values
func (eduList *EduLinkedList) createLinkedList(values []int) {
	if len(values) == 0 {
		eduList.head = nil
		return
	}

	eduList.head = &Node{Val: values[0]} // Directly creating the first Node
	current := eduList.head
	for _, value := range values[1:] {
		current.Next = &Node{Val: value} // Directly creating Node instances
		current = current.Next
	}
	eduList.head = eduList.makeCircular(eduList.head)
}

// Function to make the list circular
func (eduList *EduLinkedList) makeCircular(head *Node) *Node {
	if head == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = head // Make it circular
	return head
}

// Display prints up to `limit` nodes of the circular linked list
func Display(head *Node, limit int) {
	if head == nil {
		fmt.Println("[]")
		return
	}

	current := head
	result := []string{}
	count := 0

	for current != nil && count < limit {
		result = append(result, fmt.Sprintf("%d", current.Val))
		current = current.Next
		count++
		if current == head {
			break
		}
	}

	result = append(result, "... (circular)")
	fmt.Println(strings.Join(result, " -> "))
}

func insert(head *Node, insertVal int) *Node {
	if head == nil {
		newNode := &Node{Val: insertVal}
		newNode.Next = newNode
		return newNode
	}

	prev := head
	curr := head.Next
	insertDone := false

	for {
		if prev.Val <= insertVal && insertVal <= curr.Val {
			insertDone = true
		} else if prev.Val > curr.Val {
			if insertVal >= prev.Val || insertVal <= curr.Val {
				insertDone = true
			}
		}

		if insertDone {
			prev.Next = &Node{Val: insertVal, Next: curr}
			return head
		}

		prev = curr
		curr = curr.Next

		if prev == head {
			break
		}
	}

	prev.Next = &Node{Val: insertVal, Next: curr}
	return head

}

func main() {
	inputs := []struct {
		list      []int
		insertVal int
	}{
		{[]int{3, 4, 1}, 2},
		{[]int{}, 1},
		{[]int{1}, 0},
		{[]int{2, 2, 2}, 3},
		{[]int{5, 1, 3}, 6},
	}

	for i, input := range inputs {
		inputLinkedList := NewEduLinkedList(input.list)
		fmt.Printf("%d. \tInput linked list: ", i+1)
		Display(inputLinkedList.head, 20)
		fmt.Printf("\tValue to insert: %d\n", input.insertVal)
		fmt.Printf("\n\tUpdated linked list:  ")
		Display(insert(inputLinkedList.head, input.insertVal), 20)
		fmt.Println("\n" + strings.Repeat("-", 100))
	}
}
