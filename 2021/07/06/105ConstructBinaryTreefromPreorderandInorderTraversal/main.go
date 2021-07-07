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

	rv := preorder[0]
	index := 0
	for i := range inorder {
		if inorder[i] == rv {
			index = i
		}
	}
	root := &TreeNode{
		Val: rv,
	}

	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return root
}
