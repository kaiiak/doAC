// https://leetcode.cn/problems/implement-queue-using-stacks/

type MyQueue struct {
	s []int
}

func Constructor() MyQueue {
	return MyQueue{s: make([]int, 0)}
}

// 将元素 x 推到队列的末尾
func (this *MyQueue) Push(x int) {
	this.s = append(this.s, x)
}

// 从队列的开头移除并返回元素
func (this *MyQueue) Pop() (val int) {
	if len(this.s) == 0 {
		return 0
	}
	if len(this.s) == 1 {
		val = this.s[0]
		this.s = this.s[0:0]
		return

	}
	val = this.s[0]
	this.s = this.s[1:len(this.s)]
	return
}

// 返回开头的元素
func (this *MyQueue) Peek() int {
	if len(this.s) > 0 {
		return this.s[0]
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	return len(this.s) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */