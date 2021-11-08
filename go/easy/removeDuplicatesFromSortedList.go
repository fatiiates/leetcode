package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	start := head
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			remove(head)
			if start.Val == head.Val {
				start = head
			}
		} else {
			head = head.Next
		}
	}
	return start
}

func remove(head *ListNode) {
	head.Next = head.Next.Next
}

func main() {
	th := ListNode{
		Val:  2,
		Next: nil,
	}

	sc := ListNode{
		Val:  1,
		Next: &th,
	}

	fr := ListNode{
		Val:  1,
		Next: &sc,
	}

	r := deleteDuplicates(&fr)

	for ; r != nil; r = r.Next {
		fmt.Println(r.Val)
	}

}
