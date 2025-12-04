package Solution

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	Head  *Node
	Tail  *Node
	Cap   int
	Store map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head
	store := make(map[int]*Node, capacity)
	return LRUCache{
		Head:  head,
		Tail:  tail,
		Cap:   capacity,
		Store: store,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Store[key]; ok {
		this.moveToHead(node)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Store[key]; ok {
		node.Val = value
		this.Store[key] = node
		this.moveToHead(node)
	} else {
		newOne := &Node{
			Key: key,
			Val: value,
		}

		this.Store[key] = newOne
		this.addToHead(newOne)
		if len(this.Store) > this.Cap {
			tail := this.removeTail()
			delete(this.Store, tail.Key)
		}
	}
}

func (this *LRUCache) moveToHead(cur *Node) {
	this.removeNode(cur)
	this.addToHead(cur)
}

func (this *LRUCache) removeTail() *Node {
	last := this.Tail.Pre
	pre := last.Pre
	pre.Next = this.Tail
	this.Tail.Pre = pre

	return last
}

func (this *LRUCache) addToHead(cur *Node) {
	head := this.Head
	next := head.Next
	head.Next = cur
	cur.Pre = head
	cur.Next = next
	next.Pre = cur
}

func (this *LRUCache) removeNode(cur *Node) {
	pre := cur.Pre
	next := cur.Next
	pre.Next = next
	next.Pre = pre
}
