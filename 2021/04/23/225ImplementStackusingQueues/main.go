package main

type MyStack struct {
	arr []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.arr = append([]int{x}, this.arr...)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.arr[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.arr) == 0
}
