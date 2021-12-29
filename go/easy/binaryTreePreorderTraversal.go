package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {

	arr := []int{}
	if root == nil {
		return []int{}
	}

	arr = append(arr, root.Val)

	if root.Left != nil {
		arr = append(arr, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		arr = append(arr, preorderTraversal(root.Right)...)
	}

	return arr
}

func preorderTraversalIterative(root *TreeNode) []int {

	arr := []int{}
	backTrackArray := []*TreeNode{}
	l := 0
	for root != nil || l > 0 {
		arr = append(arr, root.Val)
		if root.Right != nil {
			backTrackArray = append(backTrackArray, root.Right)
		}
		root = root.Left

		l = len(backTrackArray)
		if root == nil && l > 0 {
			l--
			root = backTrackArray[l]
			backTrackArray = backTrackArray[:l]
		}

	}

	return arr
}
