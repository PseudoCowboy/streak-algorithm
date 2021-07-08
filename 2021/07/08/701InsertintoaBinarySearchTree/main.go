package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	helper(root, val)
	return root
}

func helper(root *TreeNode, val int) {
	if root.Val < val {
		if root.Right != nil {
			helper(root.Right, val)
		} else {
			root.Right = &TreeNode{Val: val}
		}
	}
	if root.Val > val {
		if root.Left != nil {
			helper(root.Left, val)
		} else {
			root.Left = &TreeNode{Val: val}
		}
	}
}

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root
	for cur != nil {
		if cur.Val > val {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur.Left = &TreeNode{Val: val}
				return root
			}
		}
		if cur.Val < val {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur.Right = &TreeNode{Val: val}
				return root
			}
		}
	}
	return root
}
