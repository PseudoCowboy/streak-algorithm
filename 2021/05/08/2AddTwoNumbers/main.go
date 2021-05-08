package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	carry := 0
	val := 0
	for l1 != nil && l2 != nil {
		val = l1.Val + l2.Val + carry
		carry = val / 10
		current.Next = &ListNode{Val: val % 10}
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val = l1.Val + carry
		carry = val / 10
		current.Next = &ListNode{Val: val % 10}
		current = current.Next
		l1 = l1.Next
	}
	for l2 != nil {
		val = l2.Val + carry
		carry = val / 10
		current.Next = &ListNode{Val: val % 10}
		current = current.Next
		l2 = l2.Next
	}
	if carry == 1 {
		current.Next = &ListNode{Val: 1}
	}
	return head.Next
}
