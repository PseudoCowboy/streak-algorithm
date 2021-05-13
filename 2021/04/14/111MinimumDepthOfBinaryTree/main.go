package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left < right {
		return 1 + left
	}

	return 1 + right

}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	depth := 1
	for len(stack) != 0 {
		current := len(stack)
		for i := 0; i < current; i++ {
			if stack[i].Left == nil && stack[i].Right == nil {
				return depth
			}
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[current:]
		depth++
	}
	return depth
}
