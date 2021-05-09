package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	mid := preorder[0]
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == mid {
			index = i
			break
		}
	}
	root := &TreeNode{Val: mid}
	if index != 0 {
		root.Left = buildTree(preorder[1:1+index], inorder[:index])
	}
	if index != len(inorder)-1 {
		root.Right = buildTree(preorder[1+1+index:], inorder[index+1:])
	}

	return root
}
