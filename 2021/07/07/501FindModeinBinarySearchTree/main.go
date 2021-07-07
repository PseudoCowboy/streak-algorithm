package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	cur := root
	m := make(map[int]int)
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		m[stack[len(stack)-1].Val]++
		cur = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	ans := make([]int, 0)
	freq := 0
	for k, v := range m {
		if v > freq {
			freq = v
			ans = []int{k}
		} else if v == freq {
			ans = append(ans, k)
		}
	}

	return ans
}

func findMode1(root *TreeNode) []int {
	stack := []*TreeNode{root}
	m := make(map[int]int)
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pop != nil {
			m[pop.Val]++
			stack = append(stack, pop.Left, pop.Right)
		}
	}
	ans := make([]int, 0)
	freq := 0
	for k, v := range m {
		if v > freq {
			freq = v
			ans = []int{k}
		} else if v == freq {
			ans = append(ans, k)
		}
	}

	return ans
}
