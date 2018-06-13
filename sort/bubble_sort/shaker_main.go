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
	left := 0
	right := n - 1

	// left  以左已有序
	// right 以右已有序
	// 两个区间游标相遇则集合有序
	for left < right {
		// 从左到右，选出最大值
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--
		fmt.Println("[DEBUG right]:  ", arr)
		// 从右到左，选出最小值
		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		left++
		fmt.Println("[DEBUG left]:  ", arr)
	}

	fmt.Println("[SORTED]:  ", arr)
}
