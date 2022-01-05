package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	arr []*TreeNode
}

func createQueue() Queue {
	return Queue{
		arr: []*TreeNode{},
	}
}

func (q *Queue) push(val *TreeNode) {
	q.arr = append(q.arr, val)
}

func (q *Queue) pop() *TreeNode {
	node := q.arr[0]
	q.arr = q.arr[1:]
	return node
}

func (q *Queue) len() int {
	return len(q.arr)
}

func (q *Queue) clear() {
	q.arr = []*TreeNode{}
}

func (q *Queue) set(arr []*TreeNode) {
	q.arr = arr
}

func (q *Queue) copy() []*TreeNode {
	return q.arr
}

func (q *Queue) reverse() {
	l := len(q.arr)
	for i := 0; i < l/2; i++ {
		q.arr[i], q.arr[l-i-1] = q.arr[l-i-1], q.arr[i]
	}
}

func invertTree(root *TreeNode) *TreeNode {
	pre := createQueue()
	next := createQueue()
	r := root
	if root != nil {
		pre.push(root)
		next.push(root.Left)
		next.push(root.Right)
	}

	for i := pre.len(); i > 0; i = pre.len() {
		if next.len() == 0 {
			break
		}
		next.reverse()
		for ; i > 0; i-- {
			root = pre.pop()
			if t := next.pop(); t == nil {
				root.Left = nil
			} else {
				root.Left = t
			}
			if t := next.pop(); t == nil {
				root.Right = nil
			} else {
				root.Right = t
			}
			
			pre.push(root.Left)
			pre.push(root.Right)
			if root.Right != nil {
				next.push(root.Right.Left)
				next.push(root.Right.Right)
			}
			if root.Left != nil {
				next.push(root.Left.Left)
				next.push(root.Left.Right)
			}
		}
	}

	return r
}
