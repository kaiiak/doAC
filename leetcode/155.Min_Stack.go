// https://leetcode.com/problems/min-stack/

type MinStack struct {
    stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{stack: make([]int, 0)}
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    min := this.stack[0]
    for i := 0; i < len(this.stack); i++ {
        if this.stack[i] < min {
            min = this.stack[i]
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */