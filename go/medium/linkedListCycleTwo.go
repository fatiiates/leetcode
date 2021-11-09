package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	m := make(map[*ListNode]bool)

	for ; head != nil; head = head.Next {

		if _, ok := m[head]; ok {

			return head
		} else {
			m[head] = true
		}

	}
	return nil
}

func main() {

	th := ListNode{
		Val:  2,
		Next: nil,
	}

	sc := ListNode{
		Val:  3,
		Next: &th,
	}

	fr := ListNode{
		Val:  1,
		Next: &sc,
	}
	th.Next = &fr
	fmt.Println(detectCycle(&fr))

}
