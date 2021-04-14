package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var left, right bool
	if root.Left != nil {
		left = hasPathSum(root.Left, targetSum-root.Val)
	}
	if root.Right != nil {
		right = hasPathSum(root.Right, targetSum-root.Val)
	}

	if root.Right == nil && root.Left == nil {
		return targetSum == root.Val
	}

	return left || right
}
