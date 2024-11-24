package main

// Definiton of a binary tree node class
// type EduTreeNode struct {
// 	data   int
// 	left   *EduTreeNode
// 	right  *EduTreeNode
// 	parent *EduTreeNode
// }

import (
	"fmt"
)

func lowestCommonAncestor(p, q *EduTreeNode) *EduTreeNode {

	left, right := p, q

	for left != right {
		fmt.Println(left.data, right.data)

		if left.parent == nil {
			left = q
		} else {
			left = left.parent
		}
		if right.parent == nil {
			right = p
		} else {
			right = right.parent
		}

	}
	return left
}
