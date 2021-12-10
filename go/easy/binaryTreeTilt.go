package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var total int

func findTilt(root *TreeNode) int {

	if root == nil {
		return 0
	}

	total = 0

	dfs(root)

	return total
}

func dfs(n *TreeNode) int {
	lval := 0
	rval := 0

	if n.Left != nil {
		lval = n.Left.Val
		lval += dfs(n.Left)
	}
	if n.Right != nil {
		rval = n.Right.Val
		rval += dfs(n.Right)
	}

	total += abs(lval - rval)
	return lval + rval
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
