package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next.Next.Next
		temp1 := cur.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = temp1
		cur.Next.Next.Next = temp
		cur = cur.Next.Next
	}

	return dummy.Next
}
