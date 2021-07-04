package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

func maxDepth1(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	ans++
	stack := make([]*TreeNode, 0)
	stack = append(stack, root.Left, root.Right)
	for len(stack) != 0 {
		ans++
		n := len(stack)
		for i := range stack {
			if stack[i] != nil {
				stack = append(stack, stack[i].Left, stack[i].Right)
			}
		}

		stack = stack[n:]
	}

	return ans - 1
}
