package minstack

// LeetCode #155.
// https://leetcode.com/problems/min-stack/

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// To track the min, we can:
// (1) Maintain a separate array which has the min value for that position of
//  the stack **and below**.
// (2) If we push a new value on, we only need to compare it to the min from
//  the head of the "min array". Then push the min of those to the head.
// (3) When we pop an element off the stack, we just need to remove the head of
//  the min array as well.

type MinStack struct {
	mins  []int
	elems []int
}

func Constructor() MinStack {
	return MinStack{elems: []int{}, mins: []int{}}
}

func (this *MinStack) Push(val int) {
	newMin := val
	if len(this.mins) > 0 && this.GetMin() < newMin {
		newMin = this.GetMin()
	}
	this.elems = append(this.elems, val)
	this.mins = append(this.mins, newMin)
}

func (this *MinStack) Pop() {
	this.elems = this.elems[:len(this.elems)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.elems[len(this.elems)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
