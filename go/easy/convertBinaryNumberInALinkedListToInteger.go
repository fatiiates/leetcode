package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	num := 0
	for ; head != nil; head = head.Next {
		num <<= 1
		num |= head.Val
	}
	return num
}
