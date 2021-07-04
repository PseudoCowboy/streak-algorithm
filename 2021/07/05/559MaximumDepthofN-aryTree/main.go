package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	ans := 0
	if root == nil {
		return ans
	}

	ans++
	stack := make([]*Node, 0)
	stack = append(stack, root.Children...)
	for len(stack) != 0 {
		ans++
		n := len(stack)
		for i := range stack {
			if stack[i] != nil {
				stack = append(stack, stack[i].Children...)
			}
		}
		stack = stack[n:]
	}

	return ans
}
