package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	arr := make([]int, 0)
	cur := root
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		arr = append(arr, stack[len(stack)-1].Val)
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]

	}
	min := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
		}
	}
	return min
}

func getMinimumDifference1(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	cur := root
	var pre *TreeNode
	min := math.MaxInt32
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		if pre != nil && min > stack[len(stack)-1].Val-pre.Val {
			min = stack[len(stack)-1].Val - pre.Val
		}
		pre = stack[len(stack)-1]
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return min
}
