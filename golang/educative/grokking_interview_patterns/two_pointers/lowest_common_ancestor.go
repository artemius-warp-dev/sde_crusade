package main

import (
	"fmt"
	"strings"
)

type EduTreeNode struct {
	data   int
	left   *EduTreeNode
	right  *EduTreeNode
	parent *EduTreeNode
}

type EduBinaryTree struct {
	root *EduTreeNode
}

func createBinaryTree(nodes []int) *EduTreeNode {
	if len(nodes) == 0 || nodes[0] == 0 {
		return nil
	}
	root := &EduTreeNode{data: nodes[0]}
	queue := []*EduTreeNode{root}
	i := 1
	for i < len(nodes) {
		curr := queue[0]
		queue = queue[1:]
		if i < len(nodes) && nodes[i] != 0 {
			curr.left = &EduTreeNode{data: nodes[i], parent: curr}
			queue = append(queue, curr.left)
		}
		i++
		if i < len(nodes) && nodes[i] != 0 {
			curr.right = &EduTreeNode{data: nodes[i], parent: curr}
			queue = append(queue, curr.right)
		}
		i++
	}
	return root
}

func (tree *EduBinaryTree) find(root *EduTreeNode, value int) *EduTreeNode {
	if root == nil {
		return nil
	}
	queue := []*EduTreeNode{root}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if currentNode.data == value {
			return currentNode
		}
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
	return nil
}

func lowestCommonAncestor(p, q *EduTreeNode) *EduTreeNode {
	ptr1, ptr2 := p, q

	for ptr1 != ptr2 {
		if ptr1.parent != nil {
			ptr1 = ptr1.parent
		} else {
			ptr1 = q
		}

		if ptr2.parent != nil {
			ptr2 = ptr2.parent
		} else {
			ptr2 = p
		}
	}

	return ptr1
}

// Driver code
func main() {
	inputTrees := [][]int{
		{100, 50, 200, 25, 75, 350},
		{100, 200, 75, 50, 25, 350},
		{350, 100, 75, 50, 200, 25},
		{100, 50, 200, 25, 75, 350},
		{25, 50, 75, 100, 200, 350},
	}
	inputNodes := [][]int{
		{25, 75},
		{50, 350},
		{100, 200},
		{50, 25},
		{350, 200},
	}

	for i := 0; i < len(inputTrees); i++ {
		tree := &EduBinaryTree{root: createBinaryTree(inputTrees[i])}
		fmt.Printf("%d.\tBinary tree:\n", i+1)
		//displayTree(tree.root)
		fmt.Printf("\n\tp = %d\n", inputNodes[i][0])
		fmt.Printf("\tq = %d\n", inputNodes[i][1])
		p := tree.find(tree.root, inputNodes[i][0])
		q := tree.find(tree.root, inputNodes[i][1])
		lca := lowestCommonAncestor(p, q)
		fmt.Printf("\n\tLowest common ancestor: %d\n", lca.data)
		fmt.Println(strings.Repeat("-", 100))
	}
}
