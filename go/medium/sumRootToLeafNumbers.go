package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var total int

func sumNumbers(root *TreeNode) int {
	total = 0
	travel(root, []int{})
	return total
}

func travel(n *TreeNode, sumArr []int) bool {

	sumArr = append(sumArr, n.Val)

	if n.Left == nil && n.Right == nil {
		sum(sumArr)
		return true
	}
	if n.Left != nil {
		travel(n.Left, sumArr)
	}

	if n.Right != nil {
		travel(n.Right, sumArr)
	}
	return true
}

func sum(sumArr []int) {

	sum := 0
	for i := len(sumArr) - 1; i >= 0; i-- {
		sum += int(math.Pow10(i)) * sumArr[len(sumArr)-1-i]
	}
	total += sum
}

func main() {

	left := TreeNode{
		Val:   2,
		Right: nil,
		Left:  nil,
	}

	right := TreeNode{
		Val:   3,
		Right: nil,
		Left:  nil,
	}

	root := TreeNode{
		Val:   1,
		Right: &right,
		Left:  &left,
	}
	fmt.Println(sumNumbers(&root))
}
