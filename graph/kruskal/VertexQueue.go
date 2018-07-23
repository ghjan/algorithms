package kruskal

import (
	"sync"
	"strconv"
)

type VertexQueue struct {
	items []Vertex
	lock  sync.RWMutex
}

// 创建队列
func (q *VertexQueue) New() *VertexQueue {
	q.items = []Vertex{}
	return q
}

// 入队列
func (q *VertexQueue) Enqueue(t Vertex) {
	q.lock.Lock()
	q.items = append(q.items, t)
	q.lock.Unlock()
}

// 出队列
func (q *VertexQueue) Dequeue() *Vertex {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

// 出队列
func (q *VertexQueue) Remove() {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) > 0 {
		q.items = q.items[1:len(q.items)]
	}
}

//Peek 获取队列的第一个元素，不移除 和Front函数同义
func (q *VertexQueue) Peek() *Vertex {
	return q.Front()
}

// 获取队列的第一个元素，不移除
func (q *VertexQueue) Front() *Vertex {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

// 判空
func (q *VertexQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// 获取队列的长度
func (q *VertexQueue) Size() int {
	return len(q.items)
}
func (q *VertexQueue) Sort() {
	q.items = VertexQuickSort(q.items)
}

func VertexQuickSort(arr []Vertex) []Vertex {

	n := len(arr)
	// 递归结束条件
	if n <= 1 {
		temp := make([]Vertex, n)
		copy(temp, arr)
		return temp
	}

	// 使用第一个元素的label作为基准值 转换为整数
	pivot, _ := strconv.Atoi(arr[0].Label)

	// 小元素 和 大元素各成一个数组
	low := make([]Vertex, 0, n)
	high := make([]Vertex, 0, n)

	// 更小的元素放 low[]
	// 更大的元素放 high[]
	for i := 1; i < n; i++ {
		current, _ := strconv.Atoi(arr[i].Label)
		if current < pivot {
			low = append(low, arr[i])
		} else {
			high = append(high, arr[i])
		}
	}
	// 子区间递归快排，分治排序
	low, high = VertexQuickSort(low), VertexQuickSort(high)
	//fmt.Println("[DEBUG low]:\t", low)
	//fmt.Println("[DEBUG high]:\t", high)
	return append(append(low, arr[0]), high...)
}
