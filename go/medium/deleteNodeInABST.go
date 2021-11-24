package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			rr := root.Right
			for rr.Left != nil {
				rr = rr.Left
			}
			root.Val = rr.Val
			root.Right = deleteNode(root.Right, rr.Val)
			return root
		}
	}
}
