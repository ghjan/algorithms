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
	end := n - 1
	for end > 0 {
		cur := 0
		for j := 0; j < end; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			cur = j
		}
		fmt.Println("[DEBUG]:  ", arr)
		end = cur
	}

	fmt.Println("[SORTED]:  ", arr)
}
