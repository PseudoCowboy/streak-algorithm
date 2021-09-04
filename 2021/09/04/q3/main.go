package main

//-1,0,1

func main() {
	Constructor([]int{-1, 0, 1})
}

type LockingTree struct {
	lock   []int
	parent []int
	tree   map[int][]int
}

func Constructor(parent []int) LockingTree {
	lock := make([]int, len(parent))
	tree := make(map[int][]int)
	tree[0] = make([]int, 0)
	for i := 1; i < len(parent); i++ {
		p := parent[i]
		tree[p] = append(tree[p], i)
		for parent[p] != -1 {
			p = parent[p]
			tree[p] = append(tree[p], i)
		}
	}
	return LockingTree{
		lock:   lock,
		parent: parent,
		tree:   tree,
	}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.lock[num] == 0 {
		this.lock[num] = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.lock[num] == user {
		this.lock[num] = 0
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	if this.lock[num] != 0 {
		return false
	}
	check := num
	for this.parent[check] != -1 {
		if this.lock[this.parent[check]] != 0 {
			return false
		}
		check = this.parent[check]
	}
	count := 0
	for i := range this.tree[num] {
		v := this.tree[num][i]
		if this.lock[v] != 0 {
			count++
			this.lock[v] = 0
		}
	}
	if count > 0 {
		this.lock[num] = user
		return true
	}
	return false
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
