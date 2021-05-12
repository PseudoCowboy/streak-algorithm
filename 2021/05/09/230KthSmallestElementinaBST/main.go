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

func kthSmallest1(root *TreeNode, k int) int {
	result := make([]int, 0)
	inorder2(root, &result, k)
	return result[k-1]
}

func inorder2(root *TreeNode, result *[]int, k int) {
	if len(*result) == k {
		return
	}
	if root == nil {
		return
	}

	if root.Left != nil {
		inorder2(root.Left, result)
	}
	*result = append(*result, root.Val)
	if root.Right != nil {
		inorder2(root.Right, result)
	}
}

func inorder3(root *TreeNode, result *[]int) {
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		*result = append(*result, cur.Val)
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
}
