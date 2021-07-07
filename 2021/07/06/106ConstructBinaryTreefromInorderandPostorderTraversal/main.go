package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rv := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: rv,
	}

	index := 0
	for i := range inorder {
		if inorder[i] == rv {
			index = i
			break
		}
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}
