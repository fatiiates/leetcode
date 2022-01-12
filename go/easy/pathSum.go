package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SumNode struct {
	total int
	node  *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	stack := []*SumNode{}
	var pre *TreeNode = nil

	for root != nil || len(stack) > 0 {

		for root != nil {
			if len(stack) > 0 {
				stack = append(stack, &SumNode{
					total: peek(&stack).total + root.Val,
					node:  root,
				})
			} else {
				stack = append(stack, &SumNode{
					total: root.Val,
					node:  root,
				})
			}

			root = root.Left
		}

		n := pop(&stack)

		if n.node.Right != nil && pre != n.node.Right {
			stack = append(stack, n)
			root = n.node.Right
			continue
		}
		if (n.node.Left == nil && n.node.Right == nil) && n.total == targetSum {
			return true
		}
		pre = n.node
		root = nil

	}

	return false
}

func pop(s *[]*SumNode) *SumNode {
	l := len(*s) - 1
	node := (*s)[l]
	*s = (*s)[:l]
	return node
}

func peek(s *[]*SumNode) *SumNode {
	return (*s)[len(*s)-1]
}
