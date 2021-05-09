package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	result := make([]*TreeNode, 0)
	memo := make(map[string]int)
	dfs(root, &result, memo)
	return result
}

func dfs(root *TreeNode, result *[]*TreeNode, memo map[string]int) string {
	if root == nil {
		return "#"
	}

	left := dfs(root.Left, result, memo)
	right := dfs(root.Right, result, memo)

	cur := fmt.Sprintf("%s,%s,%v", left, right, root.Val)

	memo[cur]++
	if memo[cur] == 2 {
		*result = append(*result, root)
	}

	return cur
}
