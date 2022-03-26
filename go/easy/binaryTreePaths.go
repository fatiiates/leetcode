package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	findPaths(root, &res, "")
	return res
}

func findPaths(root *TreeNode, res *[]string, s string) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			(*res) = append((*res), s+fmt.Sprint(root.Val))
		} else {
			findPaths(root.Left, res, s+fmt.Sprint(root.Val)+"->")
			findPaths(root.Right, res, s+fmt.Sprint(root.Val)+"->")
		}
	}
}

