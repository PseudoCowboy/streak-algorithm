package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	if left == -1 {
		return -1
	}
	right := helper(root.Right)
	if right == -1 {
		return -1
	}

	if left-right > 1 || right-left > 1 {
		return -1
	}

	if left > right {
		return 1 + left
	}
	return 1 + right
}

// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	left := maxHelper(root)
// 	right := maxHelper(root)

// 	if left-right >= 2 || right-left >= 2 {
// 		return false
// 	}

// 	return true
// }

// func maxHelper(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	left := maxHelper(root.Left)
// 	right := maxHelper(root.Right)

// 	if left > right {
// 		return 1 + left
// 	}

// 	return 1 + right
// }
