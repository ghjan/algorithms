package main

import (
	"fmt"
	"github.com/ghjan/algorithms"
)

func main() {
	arr := algorithms.GetArr(5, 20)
	arr = []int{29, 10, 14, 37, 14}
	fmt.Println("[UNSORTED]: ", arr)

	n := len(arr)
	if n <= 1 {
		fmt.Println("[ALREADY SORTED]: ", arr)
		return
	}
	// 遍历所有元素
	for i := 1; i < n; i++ {
		// 向前找位置
		for j := i; j > 0; j-- {
			// 合适位置插入
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		fmt.Println("[DEBUG]: ", arr)
	}
	fmt.Println("[SORTED]: ", arr)
}
