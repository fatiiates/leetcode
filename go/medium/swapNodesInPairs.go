package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	h := head

	var pre *ListNode
	for ; head != nil; head = head.Next {
		if head.Next != nil {
			if pre != nil {
				pre.Next = head.Next
			} else {
				h = head.Next
			}
			pre = head
			head.Next, head.Next.Next = head.Next.Next, head
		} else {
			break
		}
	}
	return h
}

func main() {
	m := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	n := swapPairs(&m)

	for ; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
}

/*
[1,2,3,4,5,6,7,8]
[1]
[1,2,3,4,5,6,7,8,9]
*/
