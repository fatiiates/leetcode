package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	cur := &ListNode{
		Val:  0,
		Next: nil,
	}
	for head != nil {

		node := head
		head = head.Next
		insert(cur, node)
	}

	return cur.Next
}

func insert(head *ListNode, node *ListNode) {
	for head.Next != nil && head.Next.Val < node.Val {
		head = head.Next
	}
	node.Next = head.Next
	head.Next = node
}
 
ListNode* insertionSortList(ListNode* head) {
	ListNode* new_head = new ListNode(0);

	while (head) {
		ListNode* node = head;
		head = head->next;
		insert(new_head, node);
	}
	
	return new_head->next;
}