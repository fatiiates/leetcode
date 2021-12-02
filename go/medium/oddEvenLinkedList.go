package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head

}

func main() {
	th := ListNode{
		Val:  3,
		Next: nil,
	}

	sc := ListNode{
		Val:  2,
		Next: &th,
	}

	fr := ListNode{
		Val:  1,
		Next: &sc,
	}

	t := oddEvenList(&fr)
	for ; t != nil; t = t.Next {
		fmt.Println(t.Val)
	}
}
