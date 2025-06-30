type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}
type LRUCache struct {
	head  *Node
	limit int
	l1    map[int](*Node)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{head: nil, limit: capacity, l1: make(map[int](*Node))}
}

func (this *LRUCache) add(h *Node) {
	
    if this.head == nil {
		this.head = h
		this.head.Pre = h
		return
	}

	this.head.Pre.Next = h
	h.Pre = this.head.Pre
	this.head.Pre = h
}

func (this *LRUCache) del(h *Node) {
	if this.head == nil {
		return
	}

	if h == this.head {
		this.head = this.head.Next
		if this.head != nil {
			this.head.Pre = this.head.Pre.Pre
		}
		return
	}

	if h == this.head.Pre {
		this.head.Pre.Pre.Next = nil
		this.head.Pre = this.head.Pre.Pre
		return
	}

	h.Next.Pre = h.Pre
	h.Pre.Next = h.Next
}

func (this *LRUCache) Get(key int) int {

	if ec, check := this.l1[key]; check {
		h1 := &Node{Val: ec.Val, Key: ec.Key}
		this.del(ec)
		this.add(h1)
		this.l1[key] = h1
		return h1.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	
    if _, check := this.l1[key]; check {
		this.del(this.l1[key])
	} else if len(this.l1) == this.limit {
		delete(this.l1, this.head.Key)
		this.del(this.head)
	}

	h1 := &Node{Val: value, Key: key}
	this.add(h1)
	this.l1[key] = h1
	h1 = this.head
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */