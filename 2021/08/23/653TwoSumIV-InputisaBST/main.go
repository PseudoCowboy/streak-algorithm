package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	m := make(map[int]bool)
	for len(stack) != 0 {
		size := len(stack)
		for i := range stack {
			if m[k-stack[i].Val] {
				return true
			}
			m[stack[i].Val] = true
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[size:]
	}

	return false
}
