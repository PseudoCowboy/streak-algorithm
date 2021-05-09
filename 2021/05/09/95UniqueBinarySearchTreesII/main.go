package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if start > end {
		result = append(result, nil)
		return result
	}

	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				result = append(result, root)
			}
		}
	}

	return result

}
