package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 暴力
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, subLeft, right, subRight := 0, 0, 0, 0
	if root.Left != nil {
		left = rob(root.Left)
		subLeft = rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		right = rob(root.Right)
		subRight = rob(root.Right.Left) + rob(root.Right.Right)
	}

	return max(left+right, root.Val+subLeft+subRight)
}

func rob1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r, nr := helper(root)
	return max(r, nr)
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	left, norobLeft := helper(root.Left)
	right, norobRight := helper(root.Right)

	return root.Val + norobLeft + norobRight, max(left, norobLeft) + max(right, norobRight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
