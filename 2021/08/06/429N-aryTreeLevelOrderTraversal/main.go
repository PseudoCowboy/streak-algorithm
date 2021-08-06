package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) != 0 {
		temp := make([]int, 0)
		size := len(q)
		for i := range q {
			temp = append(temp, q[i].Val)
			q = append(q, q[i].Children...)
		}
		q = q[size:]
		ans = append(ans, temp)
	}

	return ans
}
