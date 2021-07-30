package main

type MapSum struct {
	val      int
	children map[rune]*MapSum
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{val: 0, children: map[rune]*MapSum{}}
}

func (this *MapSum) Insert(key string, val int) {
	parent := this
	for _, v := range key {
		if child, ok := parent.children[v]; ok {
			parent = child
		} else {
			newChild := &MapSum{val: 0, children: map[rune]*MapSum{}}
			parent.children[v] = newChild
			parent = newChild
		}
	}
	parent.val = val
}

func (this *MapSum) Sum(prefix string) int {
	parent := this
	for _, v := range prefix {
		if child, ok := parent.children[v]; ok {
			parent = child
		} else {
			return 0
		}
	}
	q := make([]*MapSum, 0)
	sum := parent.val
	for _, v := range parent.children {
		q = append(q, v)
	}
	for len(q) > 0 {
		n := len(q)
		for i := range q {
			sum += q[i].val
			for _, v := range q[i].children {
				q = append(q, v)
			}
		}
		q = q[n:]
	}
	return sum
}
