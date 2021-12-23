package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var total int

func rangeSumBST(root *TreeNode, low int, high int) int {

	total = 0
	if root != nil && (root.Val >= low && root.Val <= high) {
		total += root.Val
	}
	if root.Left != nil {
		total += rangeSumBST(root.Left, low, high)
	}
	if root.Right != nil {
		total += rangeSumBST(root.Right, low, high)
	}

	return total
}
