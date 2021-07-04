package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		if root.Left.Val != root.Right.Val {
			return false
		} else {
			return isSymmetric(&TreeNode{Left: root.Left.Left, Right: root.Right.Right}) && isSymmetric(&TreeNode{Left: root.Left.Right, Right: root.Right.Left})
		}
	}

	return false
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root.Left, root.Right)
	for len(stack) != 0 {
		left, right := stack[0], stack[1]
		stack = stack[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		stack = append(stack, left.Left, right.Right, left.Right, right.Left)
	}

	return true
}
