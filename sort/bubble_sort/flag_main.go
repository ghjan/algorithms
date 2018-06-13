package main

import (
	"algorithms"
	"fmt"
)

func main() {
	arr := algorithms.GetArr(5, 20)
	//arr = []int{29, 10, 14, 37, 14}
	fmt.Println("[UNSORTED]:  ", arr)

	n := len(arr)
	// 遍历所有元素
	var isSorted bool
	for i := 0; i < n-1; i++ {
		isSorted = true
		for j := 0; j < n-i-1; j++ {
			// 左元素 > 右元素
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 发生交换则还未排序完毕
				isSorted = false
			}
			fmt.Println("[DEBUG]:  ", arr)
		}
		if isSorted {
			break
		}
	}

	fmt.Println("[SORTED]:  ", arr)
}
