package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	p := len(nums) / 2
	head := &TreeNode{
		Val:   nums[p],
		Left:  nil,
		Right: nil,
	}

	hl := head
	hr := head

	for i := p - 1; i >= 0; i-- {
		hl.Left = &TreeNode{
			Val:   nums[i],
			Left:  nil,
			Right: nil,
		}
		hl = hl.Left
	}

	for i := p + 1; i < len(nums); i++ {
		hr.Right = &TreeNode{
			Val:   nums[i],
			Left:  nil,
			Right: nil,
		}
		hr = hr.Right
	}

	return head
}
