package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	arr := []*ListNode{}

	for ; head != nil; head = head.Next {
		arr = append(arr, head)
	}

	l := len(arr) / 2

	return arr[l]

}

func main() {

	th1 := ListNode{
		Val:  5,
		Next: nil,
	}

	th := ListNode{
		Val:  2,
		Next: &th1,
	}

	sc := ListNode{
		Val:  1,
		Next: &th,
	}

	fr := ListNode{
		Val:  3,
		Next: &sc,
	}
	fmt.Println(middleNode(&fr))

}
