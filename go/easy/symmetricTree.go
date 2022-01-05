package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	arr := []*TreeNode{}
	if root != nil {
		arr = append(arr, root.Left, root.Right)
	}

	for i := len(arr); i > 0; i = len(arr) {

		p := i / 2
		if i%2 != 0 {
			for j := 0; j < p; j++ {
				if arr[j].Val != arr[i-j-1].Val {
					return false
				}
			}
		}

		for ; i > 0; i-- {
			root = pop(&arr)
			if root != nil {
				arr = append(arr, root.Left, root.Right)
			}
		}
	}

	return true
}

func pop(s *[]*TreeNode) *TreeNode {
	node := (*s)[0]
	*s = (*s)[1:]
	return node
}
