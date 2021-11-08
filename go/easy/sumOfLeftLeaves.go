package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {

	total := 0
	fmt.Println(root)
	if root.Left != nil && root.Left.Right == nil && root.Left.Left == nil {
		total += root.Left.Val
	} else if root.Left != nil {
		total += sumOfLeftLeaves(root.Left)

	}

	if root.Right != nil {
		total += sumOfLeftLeaves(root.Right)
	}

	return total
}

func main() {
	lsr := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	lsl := TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}

	lsr1 := TreeNode{
		Val:   20,
		Left:  &lsl,
		Right: &lsr,
	}
	lsl1 := TreeNode{
		Val:   9,
		Left:  &lsl,
		Right: &lsr,
	}
	r := TreeNode{
		Val:   3,
		Left:  &lsl1,
		Right: &lsr1,
	}

	fmt.Println(sumOfLeftLeaves(&r))

}
