package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	arr := []*ListNode{}

	for h := head; h != nil; h = h.Next {
		arr = append(arr, h)
	}

	l := len(arr)
	i := l - n

	if l == 1 {
		return nil
	}

	if i == 0 {
		return head.Next
	}

	if n == 1 {
		arr[i-1].Next = nil
		return head
	}

	*(arr[i-1].Next) = *(arr[i].Next)
	return head

}

func main() {

	fr := ListNode{
		Val:  3,
		Next: nil,
	}

	e := removeNthFromEnd(&fr, 1)

	for ; e != nil; e = e.Next {
		println(e.Val)
	}

}
