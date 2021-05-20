package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		stackLen := len(stack)
		currentArr := make([]int, stackLen)
		for i := 0; i < stackLen; i++ {
			currentArr[i] = stack[i].Val
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		result = append(result, currentArr)
		stack = stack[stackLen:]
	}
	return result
}
