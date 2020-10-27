type CQueue struct {
	Ele []int
}

func Constructor() CQueue {
	return CQueue{Ele: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.Ele = append(this.Ele, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Ele) == 0 {
		return -1
	}

	head := this.Ele[0]
	this.Ele = this.Ele[1:]
	return head
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */