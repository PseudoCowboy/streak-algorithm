package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	first := head
	second := head
	count := 0
	for i := 0; i < n; i++ {
		if first.Next == nil && count < n-1 {
			return head
		}
		first = first.Next
		count++
	}
	if first == nil {
		return head.Next
	}
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return head
}
