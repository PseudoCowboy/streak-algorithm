package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	helper(root)
	ans := 0
	dfs(root, &ans, root.Val)
	return ans % 1000000007
}

func dfs(root *TreeNode, ans *int, sum int) {
	if root.Left != nil {
		if current := root.Left.Val * (sum - root.Left.Val); current > *ans {
			*ans = current
		}
		dfs(root.Left, ans, sum)
	}
	if root.Right != nil {
		if current := root.Right.Val * (sum - root.Right.Val); current > *ans {
			*ans = current
		}
		dfs(root.Right, ans, sum)
	}
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	root.Val += helper(root.Left)
	root.Val += helper(root.Right)

	return root.Val
}

func maxProduct1(root *TreeNode) int {
	sum := 0
	dfs1(root, &sum)
	max := 0
	count(root, sum, &max)
	return max % 1000000007
}

func count(root *TreeNode, sum int, max *int) int {
	if root == nil {
		return 0
	}

	left := count(root.Left, sum, max)
	right := count(root.Right, sum, max)
	if (sum-left)*left > *max {
		*max = (sum - left) * left
	}
	if (sum-right)*right > *max {
		*max = (sum - right) * right
	}
	return left + right + root.Val
}

func dfs1(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	*sum += root.Val
	dfs1(root.Left, sum)
	dfs1(root.Right, sum)
}
