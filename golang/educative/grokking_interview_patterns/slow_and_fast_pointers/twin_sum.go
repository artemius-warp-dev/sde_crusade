package main

import (
	. "golang-test-code/ds_v1/LinkedList"
)

// Definition of a binary tree node
// type LinkedListNode[T any] struct {
//     Data T
//     Next *LinkedListNode[T]
// }

func twinSum(head *LinkedListNode[int]) int {

	// Replace this placeholder return statement with your code
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next == nil {
			fast = fast.Next
		} else {
			fast = fast.Next.Next
		}

	}

	var prev, next, curr *LinkedListNode[int] = nil, nil, slow

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr = head

	maxSum := 0
	sum := 0
	for prev != nil {
		sum = curr.Data + prev.Data
		if sum > maxSum {
			maxSum = sum
		}
		curr = curr.Next
		prev = prev.Next
	}

	return maxSum
}
