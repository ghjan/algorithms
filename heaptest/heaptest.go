package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0: n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{100, 16, 4, 8, 70, 2, 36, 22, 5, 12}

	fmt.Println("\nHeap:")
	heap.Init(h)

	fmt.Printf("最小值: %d\n", (*h)[0])

	//for(Pop)依次输出最小值,则相当于执行了HeapSort
	fmt.Println("\nHeap sort:")
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	//增加一个新值,然后输出看看
	fmt.Println("\nPush(h, 3),然后输出堆看看:")
	heap.Push(h, 3)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	fmt.Println("\n使用sort.Sort排序:")
	h2 := IntHeap{100, 16, 4, 8, 70, 2, 36, 22, 5, 12}
	sort.Sort(h2)
	for _, v := range h2 {
		fmt.Printf("%d ", v)
	}
}
