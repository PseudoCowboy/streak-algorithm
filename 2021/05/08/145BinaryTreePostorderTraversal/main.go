package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	if root.Left != nil {
		left := postorderTraversal(root.Left)
		result = append(result, left...)
	}

	if root.Right != nil {
		right := postorderTraversal(root.Right)
		result = append(result, right...)
	}

	result = append(result, root.Val)
	return result
}

func postorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	cur := root
	stack := make([]*TreeNode, 0)

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			result = append(result, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1].Left
		stack = stack[:len(stack)-1]
	}
	return reverse(result)
}

func reverse(arr []int) []int {
	start, end := 0, len(arr)-1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
