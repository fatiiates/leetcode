package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if (p.Val <= root.Val && q.Val >= root.Val) || (p.Val >= root.Val && q.Val <= root.Val) {
		return root
	} else if p.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
