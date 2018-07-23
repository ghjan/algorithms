package main

import (
	"fmt"

	"github.com/ghjan/algorithms"
	"github.com/ghjan/algorithms/sort"
)

func main() {

	arr := algorithms.GetArr(100, 1000) //数组由5个随机数组成，每个随机数：20以内的非负整数
	//arr = []int{27, 38, 12, 39, 27, 16}
	fmt.Println("[UNSORTED]:\t", arr)
	fmt.Println("[SORTED]:\t", sort.IntQuickSort(arr))
}
