package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	m := make(map[*ListNode]bool)

	for ; head != nil; head = head.Next {

		if _, ok := m[head]; ok {

			return true
		} else {
			m[head] = true
		}

	}
	return false
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
	th.Next = &fr
	fmt.Println(hasCycle(&fr))

}
