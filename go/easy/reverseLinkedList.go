package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	return reverse(nil, head)

}

func reverse(next *ListNode, pre *ListNode) *ListNode {

	tmp := pre.Next

	pre.Next = next

	if tmp == nil {
		return pre
	}
	return reverse(pre, tmp)
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
		Val:  4,
		Next: &sc,
	}

	for h := &fr; h != nil; h = h.Next {
		fmt.Println(h)
	}

	n := reverseList(&fr)
	fmt.Println("tttt")
	for h := n; h != nil; h = h.Next {
		fmt.Println(h)
	}
}
