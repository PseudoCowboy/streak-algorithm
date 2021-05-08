package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	if root.Left != nil {
		left := inorderTraversal(root.Left)
		result = append(result, left...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		right := inorderTraversal(root.Right)
		result = append(result, right...)
	}

	return result
}

func inorderTraversal1(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		result = append(result, cur.Val)
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}

	return result
}
