package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	h := head
	if h == nil {
		return h
	}

	for h.Next != nil {

		if head == h && h.Val == val {
			head = h.Next
			h = h.Next
		} else if h.Next.Val == val {
			if h.Next.Next != nil {
				*(h.Next) = *(h.Next.Next)
			} else {
				h.Next = nil
				if h.Next == nil {
					break
				}
				h = h.Next
			}
		} else {
			h = h.Next
		}

	}
	if head != nil && head.Val == val {
		head = head.Next
	}
	return head
}

func main() {

	fr5 := ListNode{
		Val:  6,
		Next: nil,
	}
	fr4 := ListNode{
		Val:  6,
		Next: &fr5,
	}
	fr3 := ListNode{
		Val:  6,
		Next: &fr4,
	}
	fr2 := ListNode{
		Val:  5,
		Next: &fr3,
	}
	fr1 := ListNode{
		Val:  6,
		Next: &fr2,
	}
	fr := ListNode{
		Val:  6,
		Next: &fr1,
	}

	e := removeElements(&fr, 6)

	for ; e != nil; e = e.Next {
		println(e.Val)
	}

}
