package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := pruneTree(root.Left)
	right := pruneTree(root.Right)

	root.Left = left
	root.Right = right
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root

}
