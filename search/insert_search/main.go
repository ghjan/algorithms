package main

import (
	"fmt"
	"github.com/ghjan/algorithms"
)

func main() {
	arr := algorithms.GetRange(1, 10000)
	initLeft := 0
	initRight := len(arr) - 1
	fmt.Println(insertSearch(arr, 233, initLeft, initRight)) // 2
}

//
// 插值查找
// 相比二分查找，修正中点的取值为自适应
//
func insertSearch(arr []int, v, left, right int) int {
	// 值需转为浮点数计算
	// 否则 leftV/allV 是取整，与顺序查找无异
	leftV := float64(v - arr[left])
	allV := float64(arr[right] - arr[left])
	diff := float64(right - left)
	// 自适应的中点取值方法
	mid := int(float64(left) + leftV/allV*diff)

	fmt.Println("[DEBUG arr[mid]]:\t", mid)

	// 找区间
	if arr[mid] == v {
		return mid
	}
	if arr[mid] > v {
		// 在左区间
		return insertSearch(arr, v, left, mid-1)
	}
	if arr[mid] < v {
		// 在右区间
		return insertSearch(arr, v, mid+1, right)
	}
	return -1
}
