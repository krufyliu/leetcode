/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */
package leetcode

type MinStack struct {
	values []int
	mins   []int
	l      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
	if this.l == 0 {
		this.mins = append(this.mins, x)
	} else {
		min := this.mins[this.l-1]
		if min > x {
			this.mins = append(this.mins, x)
		} else {
			this.mins = append(this.mins, min)
		}
	}
	this.l++
}

func (this *MinStack) Pop() {
	if len(this.values) == 0 {
		return
	}

	l := this.l
	this.values = this.values[0 : l-1]
	this.mins = this.mins[0 : l-1]
	this.l--
}

func (this *MinStack) Top() int {
	if this.l == 0 {
		panic("minStack is empty")
	}

	return this.values[this.l-1]
}

func (this *MinStack) GetMin() int {
	if this.l == 0 {
		panic("minStack is empty")
	}
	l := len(this.values)
	return this.mins[l-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
