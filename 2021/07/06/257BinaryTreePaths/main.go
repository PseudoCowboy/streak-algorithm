package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ans := make([]string, 0)
	current := strconv.Itoa(root.Val)
	if root.Left != nil {
		helper(root.Left, current, &ans)
	}
	if root.Right != nil {
		helper(root.Right, current, &ans)
	}
	if root.Left == nil && root.Right == nil {
		ans = append(ans, current)
	}
	return ans
}

func helper(root *TreeNode, current string, ans *[]string) {
	current += "->" + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, current)
	}

	if root.Left != nil {
		helper(root.Left, current, ans)
	}

	if root.Right != nil {
		helper(root.Right, current, ans)
	}
}

func binaryTreePaths1(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	ans := make([]string, 0)
	left := binaryTreePaths1(root.Left)
	for i := range left {
		ans = append(ans, strconv.Itoa(root.Val)+"->"+left[i])
	}
	right := binaryTreePaths1(root.Right)
	for i := range right {
		ans = append(ans, strconv.Itoa(root.Val)+"->"+right[i])
	}

	return ans
}
