package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalIterative(root *TreeNode) []int {
	arr := []int{}
	stack := []*TreeNode{}
	var pre *TreeNode = nil

	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = pop(&stack)

		if root.Right != nil && pre != root.Right {
			stack = append(stack, root)
			root = root.Right
			continue
		}
		arr = append(arr, root.Val)
		pre = root
		root = nil
	}

	return arr
}

func pop(s *[]*TreeNode) *TreeNode {
	node := (*s)[len(*s)-1]
	*s = append([]*TreeNode{}, (*s)[:len(*s)-1]...)
	return node
}
