package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	helper(root, targetSum, []int{}, &result)
	return result
}

func helper(root *TreeNode, targetSum int, current []int, result *[][]int) {
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		current = append(current, root.Val)
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	if root.Left != nil {
		helper(root.Left, targetSum-root.Val, append(current, root.Val), result)
	}
	if root.Right != nil {
		helper(root.Right, targetSum-root.Val, append(current, root.Val), result)
	}
}
