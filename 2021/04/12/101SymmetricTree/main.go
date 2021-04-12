package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		if root.Left.Val != root.Right.Val {
			return false
		} else {
			return isSymmetric(&TreeNode{Left: root.Left.Left, Right: root.Right.Right}) && isSymmetric(&TreeNode{Left: root.Left.Right, Right: root.Right.Left})
		}
	}

	return false
}
