package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {

	m = make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	fmt.Println(m)
	return travel(0, len(inorder)-1, &postorder)
}

func travel(s int, e int, po *[]int) *TreeNode {

	if s > e {
		return nil
	}

	var val int
	*po, val = pop(*po)

	x := TreeNode{Val: val}
	p := m[x.Val]

	x.Right = travel(p+1, e, po)
	x.Left = travel(s, p-1, po)

	return &x
}

func pop(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func print(n *TreeNode) {
	if n != nil {
		fmt.Println(n.Val)
	}
	if n.Left != nil {
		print(n.Left)
	}
	if n.Right != nil {
		print(n.Right)
	}
}

func main() {
	// n := buildTree([]int{8, 9, 10, 3, 15, 20, 7}, []int{8, 10, 9, 15, 7, 20, 3})
	n := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	print(n)
}
