package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}

	return nil
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	var queue = []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		tmp := queue[length-1]
		queue = queue[:length-1]
		if tmp != nil {
			if tmp.Val == val {
				return tmp
			}
			if val > tmp.Val {
				queue = append(queue, tmp.Right)
			} else {
				queue = append(queue, tmp.Left)
			}
		}
	}
	return nil
}
