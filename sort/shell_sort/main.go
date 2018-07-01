package main

import (
	"fmt"
	"github.com/ghjan/algorithms"
)

func main() {
	arr := algorithms.GetArr(5, 20)
	//arr = []int{4, 1, 3, 0, 2}
	fmt.Println("[UNSORTED]: ", arr)

	n := len(arr)
	if n <= 1 {
		fmt.Println("[ALREADY SORTED]: ", arr)
		return
	}

	step := n / 2

	// 步长减少到 0 则排序完毕
	for step > 0 {
		fmt.Println("[DEBUG step]: ", step)

		// 遍历第一个步长区间之后的所有元素
		for i := step; i < n; i++ {
			j := i
			// 前一个元素更大则交换值
			// j >= step	// 避免向下越界
			for j >= step && arr[j-step] > arr[j] {
				arr[j-step], arr[j] = arr[j], arr[j-step]
				j -= step
			}
			fmt.Println("[DEBUG]: ", arr)
		}
		step /= 2
	}
	fmt.Println("[SORTED]: ", arr)
}
