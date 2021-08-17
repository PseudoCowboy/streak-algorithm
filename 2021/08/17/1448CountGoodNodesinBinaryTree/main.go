package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	count := 1
	dfs(root.Left, root.Val, &count)
	dfs(root.Right, root.Val, &count)
	return count
}

func dfs(root *TreeNode, value int, count *int) {
	if root == nil {
		return
	}

	if root.Val >= value {
		*count++
		value = root.Val
	}
	dfs(root.Left, value, count)
	dfs(root.Right, value, count)
}
