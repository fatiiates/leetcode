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

func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}
	q := createQueue()
	if root != nil {
	    q.push(root)
    }
	for i := q.len(); i > 0; i = q.len() {
		values := []int{}
		for ; i > 0; i-- {
			root = q.pop()
			values = append(values, root.Val)
			if root.Left != nil {
				q.push(root.Left)
			}
			if root.Right != nil {
				q.push(root.Right)
			}
		}
		arr = append(arr, values)
	}

	return arr
}
