package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		if root.Right != nil && root.Left == nil {
			return root.Right
		}
		if root.Right != nil && root.Left != nil {
			node := getMinNode(root.Right)
			root.Val = node.Val
			root.Right = deleteNode(root.Right, node.Val)
		}
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}

	return node
}
