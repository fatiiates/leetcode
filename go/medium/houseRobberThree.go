package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var bt map[*TreeNode]int = make(map[*TreeNode]int)

func rob(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if _, ok := bt[root]; ok {
		return bt[root]
	}

	dontRob := rob(root.Left) + rob(root.Right)
	robRoot := root.Val

	if root.Left != nil {
		robRoot += rob(root.Left.Left) + rob(root.Left.Right)
	}

	if root.Right != nil {
		robRoot += rob(root.Right.Left) + rob(root.Right.Right)
	}

	bt[root] = max(dontRob, robRoot)

	return bt[root]

}

func max(n1 int, n2 int) int {
	if n1 < n2 {
		return n2
	}
	return n1
}
