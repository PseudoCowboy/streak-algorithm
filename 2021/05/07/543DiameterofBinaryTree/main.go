package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var result float64
	dfs(root, &result)
	return int(result)
}

func dfs(root *TreeNode, result *float64) float64 {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, result)
	r := dfs(root.Right, result)
	*result = math.Max(*result, l+r)
	return math.Max(l, r) + 1
}
