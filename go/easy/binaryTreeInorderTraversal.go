package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(t *TreeNode) []int {
	var stack []*TreeNode
	var arr []int

	var node *TreeNode = t

	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = push(&stack, node)
			node = node.Left
		}
		node = pop(&stack)
		arr = append(arr, node.Val)
		node = node.Right
	}
	return arr
}

func push(s *[]*TreeNode, t *TreeNode) []*TreeNode {
	return append(*s, t)
}

func pop(s *[]*TreeNode) *TreeNode {
	node := (*s)[len(*s)-1]
	*s = append([]*TreeNode{}, (*s)[:len(*s)-1]...)
	return node
}

func main() {

	lsl := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	lsr := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	r := TreeNode{
		Val:   1,
		Left:  &lsl,
		Right: &lsr,
	}

	fmt.Println(inorderTraversal(&r))

}
