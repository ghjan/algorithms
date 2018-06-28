package queue

import (
	"container/heap"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"testing"
	"fmt"
)

func equal(t *testing.T, act, exp interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n",
			filepath.Base(file), line, exp, act)
		t.FailNow()
	}
}

func TestPriorityQueue(t *testing.T) {
	c := 100
	pq := NewPriorityQueue(c)

	fmt.Println("----先压入低优先级-----")
	//先压入低优先级
	for i := 0; i < c; i++ {
		heap.Push(&pq, &QItem{Value: i, Priority: i})
	}
	equal(t, cap(pq), c*2)
	equal(t, pq.Len(), c*2)

	for i := 0; i < c; i++ {
		item := heap.Pop(&pq)
		equal(t, item.(*QItem).Value.(int), i)
	}
	equal(t, pq.Len(), c)

	pq2 := NewPriorityQueue(c)

	fmt.Println("----先压入高优先级-----")
	//先压入高优先级
	for i := c - 1; i >= 0; i-- {
		heap.Push(&pq2, &QItem{Value: i, Priority: i})
	}
	equal(t, cap(pq2), c*2)
	equal(t, pq2.Len(), c*2)

	for i := 0; i < c; i++ {
		item := heap.Pop(&pq2)
		equal(t, item.(*QItem).Value.(int), i)
	}
	equal(t, pq2.Len(), c)

}

func TestUnsortedInsert(t *testing.T) {
	c := 100
	pq := NewPriorityQueue(c)
	ints := make([]int, 0, c)

	for i := 0; i < c; i++ {
		v := rand.Int()
		ints = append(ints, v)
		heap.Push(&pq, &QItem{Value: i, Priority: v})
	}
	equal(t, pq.Len(), 2*c)
	equal(t, cap(pq), 2*c)

	//对ints里面保存的优先级进行排序 小-》大
	sort.Sort(sort.IntSlice(ints)) //sort.IntSlice 完成了sort.interface

	for i := 0; i < c; i++ {
		item, _ := pq.PeekAndShift(ints[len(ints)-1]) //每次取最小优先级
		equal(t, item.Priority, ints[i])              //验证每次取出来的item的优先级最低（最小堆）
	}
}

func TestRemove(t *testing.T) {
	c := 100
	pq := NewPriorityQueue(c)

	for i := 0; i < c; i++ {
		v := rand.Int()
		heap.Push(&pq, &QItem{Value: "test", Priority: v})
	}

	for i := 0; i < 10; i++ {
		heap.Remove(&pq, rand.Intn((c-1)-i))
	}

	lastPriority := heap.Pop(&pq).(*QItem).Priority
	for i := 0; i < (c - 10 - 1); i++ {
		item := heap.Pop(&pq)
		equal(t, lastPriority < item.(*QItem).Priority, true)
		lastPriority = item.(*QItem).Priority
	}
}
