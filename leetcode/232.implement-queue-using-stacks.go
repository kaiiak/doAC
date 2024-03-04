// https://leetcode.cn/problems/implement-queue-using-stacks/

// 请你仅使用两个栈实现先入先出队列。

type MyStack struct {
	s []int
}

func (s *MyStack) Pop() int {
	if len(s.s) == 0 {
		return 0
	}
	val := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return val
}

func NewStack() *MyStack {
	return &MyStack{s: make([]int, 0)}
}

func (s *MyStack) Push(val int) {
	s.s = append(s.s, val)
}

func (s *MyStack) Len() int {
	return len(s.s)
}

type MyQueue struct {
	in, out *MyStack
}

func Constructor() MyQueue {
	return MyQueue{in: NewStack(), out: NewStack()}
}

func (this *MyQueue) in2out() {
	if this.out.Len() == 0 {
		for this.in.Len() > 0 {
			this.out.Push(this.in.Pop())
		}
	}
}

// 将元素 x 推到队列的末尾
func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

// 从队列的开头移除并返回元素
func (this *MyQueue) Pop() (val int) {
	this.in2out()
	return this.out.Pop()
}

// 返回开头的元素
func (this *MyQueue) Peek() int {
	this.in2out()
	if this.out.Len() > 0 {
		return this.out.s[this.out.Len()-1]
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	return this.in.Len() == 0 && this.out.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */