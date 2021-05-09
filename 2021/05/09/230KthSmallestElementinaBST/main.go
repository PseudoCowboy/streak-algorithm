package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0)
	inorder(root, &arr)
	return arr[k-1]
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

func kthSmallest1(root *TreeNode, k int) int {
	index, res := 0, 0
	inorder1(root, k, &index, &res)
	return res
}

func inorder1(root *TreeNode, k int, index *int, res *int) {
	if root == nil {
		return
	}

	inorder1(root.Left, k, index, res)
	*index++
	if *index == k {
		*res = root.Val
		return
	}
	inorder1(root.Right, k, index, res)
}
