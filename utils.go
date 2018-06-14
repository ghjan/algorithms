package algorithms

import (
	"math/rand"
	"time"
)

//
// 获取 n 个 [0, max] 元素组成的数组
//
func GetArr(n, max int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max + 1)
	}
	return arr
}
