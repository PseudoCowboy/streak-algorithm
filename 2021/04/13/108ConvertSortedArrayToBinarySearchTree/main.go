package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}

	fmt.Println(len(arr)/2, arr[len(arr)/2+1:])
	fmt.Println(arr[:len(arr)/2])

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	node := &TreeNode{}
	node.Val = nums[len(nums)/2]
	if len(nums)/2 == 0 {
		return node
	}

	node.Left = sortedArrayToBST(nums[:len(nums)/2])
	node.Right = sortedArrayToBST(nums[len(nums)/2+1:])

	return node
}

// func sortedArrayToBST1(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}
// 	return helper(nums, 0, len(nums)-1)
// }

// func helper(nums []int, left, right int) *TreeNode {
// 	if left > right {
// 		return nil
// 	}
// 	node := new(TreeNode)
// 	node.Val = nums[len(nums)/2]
// 	if len(nums)/2 == 0 {
// 		return node
// 	}

// 	node.Left = helper(nums, left, len(nums)/2)
// 	node.Right = helper(nums, len(nums)/2+1, right)

// 	return node
// }
