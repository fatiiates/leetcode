package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return dfs(root, -math.MaxInt32-2, math.MaxInt32+1)

}

func dfs(root *TreeNode, min, max int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	resultL := dfs(root.Left, min, root.Val)
	resultR := dfs(root.Right, root.Val, max)

	return resultL && resultR

}
