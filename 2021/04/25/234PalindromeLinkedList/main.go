package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for start, end := 0, len(arr)-1; start < end; {
		if arr[start] == arr[end] {
			start++
			end--
			continue
		} else {
			return false
		}
	}

	return true
}
