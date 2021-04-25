package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	cur := node
	for cur.Next.Next != nil {
		cur.Val = cur.Next.Val
		cur = cur.Next
	}
	cur.Val = cur.Next.Val
	cur.Next = nil
}
