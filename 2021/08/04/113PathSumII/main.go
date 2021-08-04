package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m = map[*TreeNode]*TreeNode{}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if q[i].Left == nil && q[i].Right == nil && q[i].Val == targetSum {
				ans = append(ans, getPath(q[i]))
			}
			if q[i].Left != nil {
				q[i].Left.Val += q[i].Val
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q[i].Right.Val += q[i].Val
				q = append(q, q[i].Right)
			}
		}
		q = q[size:]
	}
	return ans
}

func getPath(leaf *TreeNode) []int {
	arr := make([]int, 0)
	current := leaf
	for current != nil {
		arr = append(arr, current.Val)
		current = m[current]
	}
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i] - arr[i+1]
	}
	reverse(arr)
	return arr
}

func reverse(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
