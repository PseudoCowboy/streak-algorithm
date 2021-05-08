package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	helper(root.Left, root.Right)
	return root
}

func helper(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right

	helper(left.Left, left.Right)
	helper(left.Right, right.Left)
	helper(right.Left, right.Right)
}

func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	pre := root
	cur := &Node{}
	for pre.Left != nil {
		cur = pre
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Left
	}
	return root
}
