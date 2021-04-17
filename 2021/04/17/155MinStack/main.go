package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a[len(a)-2])
}

type MinStack struct {
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, val, val)
	} else {
		min := this.GetMin()
		this.stack = append(this.stack, val)
		if min < val {
			this.stack = append(this.stack, min)
		} else {
			this.stack = append(this.stack, val)
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-2]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-2]
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1]
}

// type MinStack struct {
// 	length          int
// 	stack, minStack []int
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
// 	return MinStack{}
// }

// func (this *MinStack) Push(val int) {
// 	this.stack = append(this.stack, val)
// 	if this.length == 0 {
// 		this.minStack = append(this.minStack, val)
// 	} else {
// 		min := this.GetMin()
// 		if val < min {
// 			this.minStack = append(this.minStack, val)
// 		} else {
// 			this.minStack = append(this.minStack, min)
// 		}
// 	}
// 	this.length++
// }

// func (this *MinStack) Pop() {
// 	this.length--
// 	this.stack = this.stack[:this.length]
// 	this.minStack = this.minStack[:this.length]
// }

// func (this *MinStack) Top() int {
// 	return this.stack[this.length-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.minStack[this.length-1]
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
