type MyStack struct {
	Ele []int
	Len int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{Ele: []int{}, Len: 0}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Ele = append(this.Ele, x)
	this.Len++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ans := this.Ele[this.Len-1]
	this.Ele = this.Ele[:this.Len-1]
	this.Len--
	return ans
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Ele[this.Len-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Len == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */