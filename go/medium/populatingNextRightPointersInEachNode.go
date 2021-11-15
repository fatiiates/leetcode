package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	if root != nil {
		travel(root)
		return root
	}
	return nil

}

func travel(n *Node) {

	if n.Left != nil {
		n.Left.Next = n.Right
		if n.Next != nil {
			n.Right.Next = n.Next.Left
		}
		travel(n.Left)
		travel(n.Right)
	}

}
