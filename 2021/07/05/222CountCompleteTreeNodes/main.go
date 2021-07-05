package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := "123"
	fmt.Println(a[:len(a)])
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	ans := 1
	for len(stack) != 0 {
		n := len(stack)
		for i := range stack {
			if stack[i].Left != nil {
				ans++
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				ans++
				stack = append(stack, stack[i].Right)
			}
		}
		stack = stack[n:]
	}

	return ans
}

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countNodes1(root.Left)
	right := countNodes1(root.Right)

	return left + right + 1
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := root.Left
	r := root.Right
	lh := 0
	rh := 0
	for l != nil {
		lh++
		l = l.Left
	}
	for r != nil {
		rh++
		r = r.Right
	}
	if lh == rh {
		return 2<<lh - 1
	}
	return countNodes2(root.Left) + countNodes2(root.Right) + 1
}
