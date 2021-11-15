package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 != nil && root2 != nil {
		travel(root1, root2)
	} else if root1 != nil {
		return root1
	} else if root2 != nil {
		return root2
	}

	return root1
}

func travel(r1 *TreeNode, r2 *TreeNode) {

	if r1 != nil && r2 != nil {
		r1.Val += r2.Val
	} else if r1 == nil {
		r1 = r2
	}

	if r1.Left != nil && r2.Left != nil {
		travel(r1.Left, r2.Left)
	} else if r1.Left == nil && r2.Left != nil {
		r1.Left = r2.Left
	}

	if r1.Right != nil && r2.Right != nil {
		travel(r1.Right, r2.Right)
	} else if r1.Right == nil && r2.Right != nil {
		r1.Right = r2.Right
	}

}
