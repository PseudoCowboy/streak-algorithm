package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	dfs(root, &sum)
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	dfs(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	dfs(root.Left, sum)

}

func convertBST1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	cur := root
	total := 0
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		total += cur.Val
		cur.Val = total
		cur = cur.Left
		stack = stack[:len(stack)-1]
	}
	return root
}
