package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	balance := balanced(root)

	if balance == -1 {
		return false
	}

	return true

}

func balanced(root *TreeNode) int {
	left := 0
	if root.Left != nil {
		h := balanced(root.Left)
		if h == -1 {
			return -1
		}
		left += h
	}

	right := 0
	if root.Right != nil {
		h := balanced(root.Right)
		if h == -1 {
			return -1
		}
		right += h
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(n, n1 int) int {
	if n > n1 {
		return n
	}
	return n1
}

func main() {

}
