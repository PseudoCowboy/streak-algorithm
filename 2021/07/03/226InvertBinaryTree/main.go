package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func invertTree1(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop != nil {
			pop.Left, pop.Right = pop.Right, pop.Left
			stack = append(stack, pop.Left, pop.Right)
		}
	}

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		pop := stack[0]
		stack = stack[1:]
		pop.Left, pop.Right = pop.Right, pop.Left
		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}
		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
	}

	return root
}
