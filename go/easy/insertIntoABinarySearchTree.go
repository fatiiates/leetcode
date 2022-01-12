package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	r := root
	if root == nil {
		return &TreeNode{
			Left: nil,
			Right: nil,
			Val: val,
		}
	}
	for root != nil {
		if val > root.Val {
			if root.Right == nil {
				root.Right = &TreeNode{
					Left: nil,
					Right: nil,
					Val: val,
				}
				break
			}
			root = root.Right
		} else {
			if root.Left == nil {
				root.Left = &TreeNode{
					Left: nil,
					Right: nil,
					Val: val,
				}
				break
			}
			root = root.Left
		}
	}
	return r
}
