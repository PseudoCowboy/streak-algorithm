package main

type MyQueue struct {
	arr []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.arr[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}
