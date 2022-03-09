package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode = &ListNode{}
	var list *ListNode = head
	for l1 != nil || l2 != nil || carry > 0 {
		total := carry
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		carry = total / 10
		node := ListNode{
			Val:  total % 10,
			Next: nil,
		}
		list.Next = &node
		list = list.Next
	}

	return head.Next
}

func main() {
	var list *ListNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	var list2 *ListNode = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	head := addTwoNumbers(list, list2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
