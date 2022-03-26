package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	return construct(&nums, 0, len(nums)-1)
}

func findMax(nums *[]int, start, end int) int {
	index := start
	for i := start+1; i <= end; i++ {
		if (*nums)[i] > (*nums)[index] {
			index = i
		}
	}
	return index
}

func construct(nums *[]int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	index := findMax(nums, start, end)

	node := &TreeNode{
		Val:   (*nums)[index],
		Left:  nil,
		Right: nil,
	}

	node.Left = construct(nums, start, index-1)
	node.Right = construct(nums, index+1, end)

	return node
}
