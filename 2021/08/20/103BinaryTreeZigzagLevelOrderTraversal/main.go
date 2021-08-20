package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	ans := make([][]int, 0)
	for len(stack) != 0 {
		size := len(stack)
		current := make([]int, 0)
		for i := 0; i < size; i++ {
			current = append(current, stack[i].Val)
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[size:]
		ans = append(ans, current)
	}
	for i := 1; i < len(ans); i += 2 {
		reverse(ans[i])
	}
	return ans
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
