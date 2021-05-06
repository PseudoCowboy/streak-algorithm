package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	arr := make([]int, 0)
	diff := math.MaxInt16
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		arr = append(arr, stack[len(stack)-1].Val)
		current = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	for i := 1; i < len(arr); i++ {
		diff = min(abs(arr[i]-arr[i-1]), diff)
	}

	return diff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
