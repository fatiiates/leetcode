package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	return find(root, k, make(map[int]bool))
}

func find(n *TreeNode, k int, m map[int]bool) bool {
	if n == nil {
		return false
	}
	if _, ok := m[n.Val]; ok {
		return true
	}
	m[n.Val] = true

	return find(n.Right, k, m) || find(n.Right, k, m)
}
