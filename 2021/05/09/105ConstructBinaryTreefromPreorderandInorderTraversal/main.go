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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	mid := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			mid = i
		}
	}

	root := &TreeNode{Val: rootVal}
	if len(inorder[:mid]) != 0 {
		root.Left = buildTree(preorder[1:1+len(inorder[:mid])], inorder[:mid])
	}
	if len(inorder[mid+1:]) != 0 {
		root.Right = buildTree(preorder[1+len(inorder[:mid]):], inorder[mid+1:])
	}

	return root
}
