// https://leetcode.com/problems/lru-cache


// LNode 单链表节点
type LNode struct {
	Key    int
	Value  int
	Next   *LNode
	Before *LNode
}

type LRUCache struct {
	Cap       int
	Value     map[int]*LNode
	firstNode *LNode
	endNode   *LNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:   capacity,
		Value: map[int]*LNode{},
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.Value[key]; ok {
		this.Move2Head(key)
		return value.Value
	}
	return -1
}

func (this *LRUCache) Out() {
	delete(this.Value, this.endNode.Key)
	// 一个节点
	if this.endNode == this.firstNode {
		this.firstNode = nil
		this.endNode = nil
		return
	}
	// 两个节点
	if this.endNode.Before == this.firstNode {
		this.firstNode.Next = nil
	}
	// 多个节点
	this.endNode.Before.Next = nil
	this.endNode = this.endNode.Before
}

func (this *LRUCache) Move2Head(key int) {
	var node *LNode
	var ok bool
	if node, ok = this.Value[key]; !ok {
		return
	}
	// 在最前,一个节点
	if this.firstNode == node {
		return
	}
	// 在最后
	if this.endNode == node {
		this.endNode = node.Before
		node.Before.Next = nil
		node.Before = nil
		node.Next = this.firstNode
		this.firstNode.Before = node
		this.firstNode = node
		return
	}
	// 在中间
	node.Before.Next = node.Next
	node.Next.Before = node.Before
	node.Next = this.firstNode
	this.firstNode.Before = node
	this.firstNode = node
}

func (this *LRUCache) Put(key, value int) {
	if _, ok := this.Value[key]; ok {
		this.Value[key].Value = value
		this.Move2Head(key)
		return
	}
	if len(this.Value) == this.Cap {
		this.Out()
	}
	nNode := &LNode{
		Key:   key,
		Value: value,
	}
	if this.firstNode != nil {
		this.firstNode.Before = nNode
		nNode.Next = this.firstNode
		this.firstNode = nNode
		this.Value[key] = nNode
	} else {
		this.Value[key] = nNode
		this.firstNode = nNode
		this.endNode = nNode
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
