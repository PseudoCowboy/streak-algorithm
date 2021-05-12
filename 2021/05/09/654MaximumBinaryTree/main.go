package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max, index := nums[0], 0

	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
			index = i
		}
	}

	root := &TreeNode{Val: max}
	if index != 0 {
		root.Left = constructMaximumBinaryTree(nums[:index])
	}
	if index != len(nums)-1 {
		root.Right = constructMaximumBinaryTree(nums[index+1:])
	}

	return root
}

func constructMaximumBinaryTree1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := 0
	maxIndex := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}
	root := &TreeNode{Val: max}
	if maxIndex != 0 {
		root.Left = constructMaximumBinaryTree1(nums[:maxIndex])
	}
	if maxIndex != len(nums)-1 {
		root.Right = constructMaximumBinaryTree1(nums[maxIndex+1:])
	}
	return root
}
