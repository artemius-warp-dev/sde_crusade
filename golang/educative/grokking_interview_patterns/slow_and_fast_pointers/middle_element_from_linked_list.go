package main

// The code in "edu_linked_list.go" can be used to create a linked list out of a list.
// The code in "linked_list_traversal.go" can be used to traverse the linked list.
// The code in "linked_list_reversal.go" can be used to reverse the linked list.

func getMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {

	slow, fast := head, head
	for {
		if fast == nil || fast.next == nil {
			return slow
		} else {
			slow = slow.next
			fast = fast.next.next
		}
	}
	return new(EduLinkedListNode)
}
