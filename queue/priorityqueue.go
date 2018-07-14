package queue

import "container/heap"

/*
优先队列（Priority Queue）：特殊的“队列”，取出元素的顺序是依照元素的优先权（关键字）大小，而不是元素进入队列的先后顺序
其实优先级队列就是将堆进行一次封装，都调用了堆的函数。
这里利用了最小堆
https://blog.csdn.net/u012233832/article/details/79634048
*/

//QItem :队列成员
type QItem struct {
	Value    interface{}
	Index    int
	Priority int
}

//PriorityQueue :优先级队列
type PriorityQueue []*QItem

//NewPriorityQueue :新建一个优先级队列
func NewPriorityQueue(cap int) PriorityQueue {
	return make(PriorityQueue, cap)
}

// 实现接口heap.Interface接口
func (pg PriorityQueue) Len() int {
	return len(pg)
}

func (pg PriorityQueue) Less(i, j int) bool {
	return pg[j] == nil || (pg[i] != nil && pg[j] != nil && pg[i].Priority < pg[j].Priority)
}

func (pg PriorityQueue) Swap(i, j int) {
	pg[i], pg[j] = pg[j], pg[i]
	if pg[i] != nil {
		pg[i].Index = i
	}
	if pg[j] != nil {
		pg[j].Index = j
	}
}

//Push :add x as element Len() 放在最后
func (pg *PriorityQueue) Push(x interface{}) {
	item := x.(*QItem)
	if item == nil {
		return
	}
	n := len(*pg)
	c := cap(*pg)
	if n+1 > c {
		npg := make(PriorityQueue, c*2)
		copy(npg, *pg)
		*pg = npg
	}
	*pg = (*pg)[0: n+1]
	item.Index = n
	(*pg)[n] = item
}

//Pop :remove and return element Len() - 1.  把最后一个拿走
func (pg *PriorityQueue) Pop() interface{} {
	n := len(*pg)
	c := cap(*pg)
	if n < (c/2) && c > 25 {
		npg := make(PriorityQueue, n, c/2)
		copy(npg, *pg)
		*pg = npg
	}
	item := (*pg)[n-1]
	if item != nil {
		item.Index = -1
	}
	*pg = (*pg)[0: n-1]
	return item
}

//PeekAndShift :获取最小值
// max表示最大优先级
func (pg *PriorityQueue) PeekAndShift(max int) (*QItem, int) {
	if pg.Len() == 0 {
		return nil, 0
	}

	item := (*pg)[0]
	if item.Priority > max {
		return nil, item.Priority - max
	}
	heap.Remove(pg, 0)
	return item, 0
}
