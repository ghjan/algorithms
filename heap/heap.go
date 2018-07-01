package heap

import (
	"errors"
	"fmt"
	"strings"
)

const (
	MAXN = 1001
	MINH = -10001
)

//Heap 最小堆
type Heap struct {
	Keys []int //堆
	Size int
}

//Create 创建一个容量为cap的最小堆
func Create(cap int) *Heap {
	keys := make([]int, cap+1)
	H := &Heap{keys, 0}
	//设置岗哨
	H.Keys[0] = MINH
	return H
}

//Insert 将X插入H,
func (H *Heap) Insert(X int) error {
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
func (H *Heap) IsFull() bool {
	return H.Size >= MAXN-1
}

//Path 返回某个节点到根节点的路径
func (H *Heap) Path(X int) string {
	path := ""

	for j := X; j >= 1; j /= 2 { /*沿根方向输出各结点*/
		path += fmt.Sprintf(" %d", H.Keys[j])
	}
	return strings.TrimLeft(path, " ")
}
