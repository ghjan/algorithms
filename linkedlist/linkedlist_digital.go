package linkedlist

import (
	"sync"
	"github.com/ghjan/algorithms/stack"
)

const (
	MaxCapacity = 100000
)

type DigitalNode struct {
	content Item
	next    int32
}

type DigitalItemLinkedList struct {
	Nodes []DigitalNode
	size  int32
	head  int32
	lock  sync.RWMutex
}

func (list *DigitalItemLinkedList) Init() {
	list.Nodes = make([]DigitalNode, MaxCapacity, MaxCapacity)
	list.size = 0
	list.head = -1
}

// 在链表结尾追加元素
func (list *DigitalItemLinkedList) AddNode(t Item, pos, next int32) {
	list.lock.Lock()
	defer list.lock.Unlock()
	isHead := list.Nodes == nil || len(list.Nodes) == 0
	if isHead { // 空链表第一次追加元素
		list.Nodes = make([]DigitalNode, MaxCapacity, MaxCapacity)
	}
	list.Nodes[pos] = DigitalNode{t, next}
	list.size = list.size + 1
}

// 检查链表是否为空
func (list *DigitalItemLinkedList) IsEmpty() bool {
	list.lock.RLock()
	defer list.lock.RUnlock()
	return list.size == 0 || list.Nodes == nil || len(list.Nodes) == 0
}

// 获取链表的长度
func (list *DigitalItemLinkedList) Size() int32 {
	list.lock.RLock()
	defer list.lock.RUnlock()
	return list.size
}

func (list *DigitalItemLinkedList) Head() int32 {
	return list.head
}

func (list *DigitalItemLinkedList) browse(operationFunc func(nodeIndex int32)) {
	for nodeIndex := list.head; nodeIndex >= 0; {
		operationFunc(nodeIndex)
		nodeIndex = list.Nodes[nodeIndex].next
	}
}
func (list *DigitalItemLinkedList) Reverse(k int32) {
	list.lock.Lock()
	defer list.lock.Unlock()
	//newHeadFinished := false
	var count int32                     //处理节点的计数
	var lastIndexPrevSection int32 = -1 //上一段的最后一个节点index，只有第一次才是小于0的
	var startIndex int32 = list.head    //开始的第一个节点Index
	for count = 0; count < list.size; count += k {
		lastIndexPrevSection = list.processReverse(k, lastIndexPrevSection, startIndex)
		startIndex = list.Nodes[lastIndexPrevSection].next
		//newHeadFinished = true
	}
}

func (list *DigitalItemLinkedList) processReverse(k, lastIndexPrevSection, startIndex int32) (int32) {
	stk := stack.ItemStack{}
	nodeIndex := startIndex
	var i int32
	for i = 0; i < k && nodeIndex >= 0; i++ { //把k个节点的index压入堆栈
		stk.Push(nodeIndex)
		nodeIndex = list.Nodes[nodeIndex].next
	}
	isFirst := true
	var prevIndex, rNodeIndex, latestIndex int32
	var count int32 = 0
	//计数 0~k-1
	latestIndex = -1 //本段最后一个节点index，下一段用得到
	//前面一段的最后一个
	prevIndex = -1
	//前置节点
	for rNodeIndex = (*stk.Pop()).(int32); count < k && !stk.IsEmpty(); rNodeIndex = (*stk.Pop()).(int32) {
		if isFirst { //某一段的第一个
			prevIndex = rNodeIndex
			if lastIndexPrevSection < 0 { //第一段的第一个
				list.head = rNodeIndex
				//newHeadFinished = true
			} else if lastIndexPrevSection > 0 { //非第一段的第一个，前面一段的最后一个链接到这个元素
				list.Nodes[lastIndexPrevSection].next = rNodeIndex
			}
			isFirst = false
		} else if prevIndex >= 0 {
			list.Nodes[prevIndex].next = rNodeIndex
		}
		if count == k-1 { //最后一个节点
			latestIndex = rNodeIndex
			break
		}
		count ++

	}
	return latestIndex
}
