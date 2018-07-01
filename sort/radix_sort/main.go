package main

import (
	"fmt"
	"github.com/ghjan/algorithms"
)

func main() {
	arr := algorithms.GetArr(5, 1000)
	//arr = []int{27, 38, 12, 101, 27, 16}
	fmt.Println("[UNSORTED]:\t", arr)
	fmt.Println("[SORTED]:\t", radixSort(arr))
}

func radixSort(arr []int) []int {
	max := getMax(arr)
	// 数组中最大值决定了循环次数，101 循环三次
	for bit := 1; max/bit > 0; bit *= 10 {
		arr = bitSort(arr, bit)
		fmt.Println("[DEBUG bit]\t", bit)
		fmt.Println("[DEBUG arr]\t", arr)
	}
	return arr
}

//
// 对指定的位进行排序
// bit 可取 1，10，100 等值
//
func bitSort(arr []int, bit int) []int {
	n := len(arr)
	// 各个位的相同的数统计到 bitCounts[] 中
	bitCounts := make([]int, 10)
	for i := 0; i < n; i++ {
		num := (arr[i] / bit) % 10
		bitCounts[num]++
	}
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1]
	}

	tmp := make([]int, 10)
	for i := n - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		tmp[bitCounts[num]-1] = arr[i]
		bitCounts[num]--
	}
	for i := 0; i < n; i++ {
		arr[i] = tmp[i]
	}
	return arr
}

// 获取数组中最大的值
func getMax(arr []int) (max int) {
	max = arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return
}
