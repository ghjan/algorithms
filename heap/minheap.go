package heap

import (
	"errors"
	"fmt"
	"strings"
)
/*
heap 堆
minheap 最小堆/小顶堆

堆是一种经过排序的完全二叉树，其中任一非终端节点的数据值均不大于（或不小于）其左子节点和右子节点的值。

堆的两个特性
结构性：用数组表示的完全二叉树；
有序性：任一结点的关键字是其子树所有结点的最大值(或最小值)
 “最大堆(MaxHeap)”,也称“大顶堆”：最大值
 “最小堆(MinHeap)”,也称“小顶堆” ：最小值

注意：从根结点到任意结点路径上结点序列的有序性！

 */

const (
	MAXN = 1001
	MINH = -10001
)

//MinHeap 最小堆
type MinHeap struct {
	Keys []int //堆
	Size int
}

//Create 创建一个容量为cap的最小堆
func Create(cap int) *MinHeap {
	keys := make([]int, cap+1)
	H := &MinHeap{keys, 0}
	//设置岗哨
	H.Keys[0] = MINH
	return H
}

//Insert 将X插入H,
func (H *MinHeap) Insert(X int) error {
	// 检查堆是否已满的代码
	if H.IsFull() {
		return errors.New("full ")
	}
	H.Size = H.Size + 1
	var i int
	for i = H.Size; H.Keys[i/2] > X; i /= 2 {
		H.Keys[i] = H.Keys[i/2]
	}
	H.Keys[i] = X
	return nil
}

//IsFull 堆是否已经满
func (H *MinHeap) IsFull() bool {
	return H.Size >= MAXN-1
}

//Path 返回某个节点到根节点的路径
func (H *MinHeap) Path(X int) string {
	path := ""

	for j := X; j >= 1; j /= 2 { /*沿根方向输出各结点*/
		path += fmt.Sprintf(" %d", H.Keys[j])
	}
	return strings.TrimLeft(path, " ")
}
