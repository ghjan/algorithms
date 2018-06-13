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
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// 左元素 > 右元素
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println("[DEBUG]:  ", arr)
		}
	}

	fmt.Println("[SORTED]:  ", arr)
}
