package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p != nil {
		p = p.Right
	}
	p.Right = right

}
