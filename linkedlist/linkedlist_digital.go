package linkedlist

import (
	"strconv"
	"strings"
	"sync"

	"github.com/ghjan/algorithms/stack"
	"os"
	"fmt"
	"bufio"
	"io"
)

const (
	//MaxCapacity 最大容量
	MaxCapacity = 100000
)

var ZeroString = []string{"00000", "0000", "000", "00", "0"}

//DigitalNode 区别于Node：next是节点的index不是真正的节点
type DigitalNode struct {
	content Item
	next    int32
}

//DigitalItemLinkedList 使用DigitalNode的单链线性链表
type DigitalItemLinkedList struct {
	nodes []*DigitalNode
	size  int32
	head  int32
	lock  sync.RWMutex
}

//Init 初始化
func (digList *DigitalItemLinkedList) Init() {
	digList.nodes = make([]*DigitalNode, MaxCapacity, MaxCapacity)
	digList.size = 0
	digList.head = -1
}

//AddNode 在链表某个位置设置Node的值
func (digList *DigitalItemLinkedList) AddNode(t Item, pos, next int32) {
	digList.lock.Lock()
	// defer digList.lock.Unlock()
	isHead := digList.nodes == nil || len(digList.nodes) == 0
	if isHead { // 空链表第一次追加元素
		digList.nodes = make([]*DigitalNode, MaxCapacity, MaxCapacity)
	}
	digList.nodes[pos] = &DigitalNode{t, next}
	digList.size ++
	digList.lock.Unlock()
}

//IsEmpty 检查链表是否为空
func (digList *DigitalItemLinkedList) IsEmpty() bool {
	digList.lock.RLock()
	defer digList.lock.RUnlock()
	return digList.size == 0 || digList.nodes == nil || len(digList.nodes) == 0
}

//Size 获取链表的长度
func (digList *DigitalItemLinkedList) Size() int32 {
	digList.lock.RLock()
	defer digList.lock.RUnlock()
	return digList.size
}

func (digList *DigitalItemLinkedList) Head() int32 {
	return digList.head
}

func (digList *DigitalItemLinkedList) Browse(operationFunc func(nodeIndex int32)) {
	for nodeIndex := digList.head; nodeIndex >= 0; {
		operationFunc(nodeIndex)
		nodeIndex = digList.nodes[nodeIndex].next
	}
}
func (digList *DigitalItemLinkedList) Reverse(N, K int32) {
	digList.lock.Lock()
	defer digList.lock.Unlock()
	var count int32 //处理节点的计数
	//var thisSectionFirstIndex int32 = -1           //这一段的第一个节点index，
	var thisSectionLatestIndex int32 = -1          //这一段的最后一个节点index，只有第一次才是小于0的
	var nextSectionStartIndex int32 = digList.head //下一段第一个节点Index
	for count = 0; count < N; count += K {
		if nextSectionStartIndex >= 0 {
			_, thisSectionLatestIndex, nextSectionStartIndex = digList.processReverse(K, thisSectionLatestIndex, nextSectionStartIndex)
		}
	}
	//digList.head = thisSectionFirstIndex
}

func (digList *DigitalItemLinkedList) processReverse(k, lastIndexPrevSection, startIndex int32) (int32, int32, int32) {
	stk := stack.ItemStack{}
	var nodeIndex int32 = startIndex
	var i int32
	isLastSection := false
	for i = 0; i < k; i++ { //把k个节点的index压入堆栈
		if nodeIndex < 0 { //已经是最后节点了
			isLastSection = true //所以也是最后一段
			break
		}
		stk.Push(nodeIndex)
		nodeIndex = digList.nodes[nodeIndex].next
	}
	isFirst := true
	var prevIndex, rNodeIndex, thisSectionFirstIndex, thisSectionLatestIndex, nextSectionStartIndex int32
	var count int32 = 0         //计数 0~k-1
	thisSectionFirstIndex = -1  //本段第一个节点index
	thisSectionLatestIndex = -1 //本段最后一个节点index，下一段用得到
	prevIndex = -1              //用于记录前置节点index
	rNodeIndex = -1             //当前处理节点index
	nextSectionStartIndex = -1  //下一段第一个节点index

	for rNodeIndex = (*stk.Pop()).(int32); count < k; {
		if rNodeIndex < 0 {
			continue
		}
		if isFirst { //某一段的第一个
			prevIndex = rNodeIndex
			if lastIndexPrevSection < 0 { //第一段的第一个
				digList.head = rNodeIndex
				//newHeadFinished = true
			} else if lastIndexPrevSection > 0 { //非第一段的第一个，前面一段的最后一个链接到这个元素
				digList.nodes[lastIndexPrevSection].next = rNodeIndex
			}
			thisSectionFirstIndex = rNodeIndex
			nextSectionStartIndex = digList.nodes[rNodeIndex].next
			isFirst = false
		} else if prevIndex >= 0 {
			digList.nodes[prevIndex].next = rNodeIndex
			prevIndex = rNodeIndex
		}
		count++
		if stk.IsEmpty() {
			break
		}
		rNodeIndex = (*stk.Pop()).(int32)
	}
	//最后一个节点
	if rNodeIndex >= 0 {
		thisSectionLatestIndex = rNodeIndex
		if isLastSection {
			digList.nodes[thisSectionLatestIndex].next = -1 //这是最后一段的最有一个节点
		}
	}
	return thisSectionFirstIndex, thisSectionLatestIndex, nextSectionStartIndex
}

func (digList *DigitalItemLinkedList) ProcessAddNode(nodeInfo string) int32 {
	nodeInfoList := strings.Split(nodeInfo, " ")
	var pos, content, next int32
	temp, _ := strconv.ParseInt(nodeInfoList[0], 10, 32)
	pos = int32(temp)
	temp, _ = strconv.ParseInt(nodeInfoList[1], 10, 32)
	content = int32(temp)
	temp, _ = strconv.ParseInt(nodeInfoList[2], 10, 32)
	next = int32(temp)
	digList.AddNode(content, pos, next)
	return pos
}

func SolveReverseLinkedList(filename string) {

	if list, N, K, err := buildDigitalItemLinkedList(filename); err == nil {
		list.Reverse(N, K)
		list.Browse(func(nodeIndex int32) {
			node := list.nodes[nodeIndex]
			fmt.Printf("%s %d %s\n", convertToAddress(nodeIndex), node.content, convertToAddress(node.next))
		})
	} else {
		fmt.Println(err)
	}
}

func buildDigitalItemLinkedList(filename string) (*DigitalItemLinkedList, int32, int32, error) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, -1, -1, nil
	}
	defer fi.Close()

	var list DigitalItemLinkedList
	br := bufio.NewReader(fi)
	var N, K int32
	index := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if index == 0 { //第一行
			headInfoLit := strings.Split(string(a), " ")
			temp, _ := strconv.ParseInt(headInfoLit[0], 10, 32)
			headIndex := int32(temp)
			temp, _ = strconv.ParseInt(headInfoLit[1], 10, 32)
			N = int32(temp)
			temp, _ = strconv.ParseInt(headInfoLit[2], 10, 32)
			K = int32(temp)
			list = DigitalItemLinkedList{}
			list.head = headIndex
		} else {
			list.ProcessAddNode(string(a))
		}
		index++
	}
	return &list, N, K, nil
}

func convertToAddress(index int32) string {
	if index < 0 {
		return strconv.Itoa(int(index))
	}
	result := strconv.Itoa(int(index))
	if len(result) < 5 {
		result = ZeroString[len(result)] + result
	}
	return result
}
