package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	ans := 0
	if helper(root, &ans) == 0 {
		ans++
	}
	return ans
}

func helper(root *TreeNode, count *int) int {
	if root == nil {
		return 2
	}
	left := helper(root.Left, count)
	right := helper(root.Right, count)
	if left == 2 && right == 2 {
		return 0
	} else if left == 0 || right == 0 {
		*count++
		return 1
	}
	return 2
}
