package Solution

import "fmt"

var dirArr = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	matrix := make([][]bool, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]bool, n)
	}

	var dfs func(int, int, map[string]bool)
	dfs = func(x, y int, visited map[string]bool) {
		key := fmt.Sprintf("%v_%v", x, y)
		if x < 0 || x > m-1 || y < 0 || y > n-1 {
			return
		}
		if visited[key] {
			return
		}

		if board[x][y] == 'O' {
			visited[key] = true
			matrix[x][y] = true
		} else {
			return
		}
		for i := 0; i < len(dirArr); i++ {
			nx, ny := x+dirArr[i][0], y+dirArr[i][1]
			dfs(nx, ny, visited)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && (i == 0 || j == 0 || i == m-1 || j == n-1) {
				visited := make(map[string]bool)
				dfs(i, j, visited)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !matrix[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{0, 0, nil, nil}
	tail := &Node{0, 0, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
	} else {
		newNode := &Node{key, value, nil, nil}
		this.cache[key] = newNode
		this.addToHead(newNode)
		if len(this.cache) > this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.key)
		}
	}
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
