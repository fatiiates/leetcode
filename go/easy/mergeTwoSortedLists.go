package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	firstNode := l2

	for {

		if l1.Val >= l2.Val {
			if l2.Next != nil && l2.Next.Val < l1.Val {
				l2 = l2.Next
				continue
			}
			newNode := ListNode{
				Val:  l1.Val,
				Next: (*l2).Next,
			}
			l2.Next = &newNode
			l2 = &newNode
			l1 = l1.Next
		} else {
			next := l1.Next
			l1.Next = l2
			l2 = l1
			firstNode = l1
			l1 = next
		}
		if l1 == nil {
			break
		}
	}
	return firstNode
}
