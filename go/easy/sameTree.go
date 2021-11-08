package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	var stack []*TreeNode
	var arrp []int
	var arrq []int

	for len(stack) > 0 || p != nil {
		for p != nil {
			stack = push(&stack, p)
			p = p.Left
		}
		p = pop(&stack)
		arrp = append(arrp, p.Val)
		p = p.Right
	}

	for len(stack) > 0 || q != nil {
		for q != nil {
			stack = push(&stack, q)
			q = q.Left
		}
		q = pop(&stack)
		arrq = append(arrq, q.Val)
		q = q.Right
	}

	if len(arrp) != len(arrq) {
		return false
	}

	for i, v := range arrp {
		if v != arrq[i] {
			return false
		}
	}

	return true

}

func push(s *[]*TreeNode, t *TreeNode) []*TreeNode {
	return append(*s, t)
}

func pop(s *[]*TreeNode) *TreeNode {
	node := (*s)[len(*s)-1]
	*s = append([]*TreeNode{}, (*s)[:len(*s)-1]...)
	return node
}
