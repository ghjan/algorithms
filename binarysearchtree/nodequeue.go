package binarysearchtree

import (
	"sync"
)

type NodeItemQueue struct {
	items []Node
	lock  sync.RWMutex
}

// 创建队列
func (q *NodeItemQueue) New() *NodeItemQueue {
	q.items = []Node{}
	return q
}

// 入队列
func (q *NodeItemQueue) Enqueue(t *Node) {
	q.lock.Lock()
	q.items = append(q.items, *t)
	q.lock.Unlock()
}

// 出队列
func (q *NodeItemQueue) Dequeue()  *Node {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

// 获取队列的第一个元素，不移除
func (q *NodeItemQueue) Front()  *Node {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

// 判空
func (q *NodeItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// 获取队列的长度
func (q *NodeItemQueue) Size() int {
	return len(q.items)
}
