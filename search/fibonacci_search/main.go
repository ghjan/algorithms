package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(fibonacciSearch(arr, 8))
}

//
// 斐波那契查找
// 在填充后的数组中进行元素的查找
//
func fibonacciSearch(arr []int, val int) int {
	arrLen := len(arr)
	fbArr := makeFibonacciArray(arr)
	fillLen := fbArr[len(fbArr)-1]

	// 取对应斐波那契数列最后一个元素的值作为填充数组的长
	fillArr := make([]int, fillLen)

	// 先填充原数组的值
	for i, v := range arr {
		fillArr[i] = v
	}

	// 如果有空余，则填充最后一个元素值
	last := arr[arrLen-1]
	for i := arrLen; i < fillLen; i++ {
		fillArr[i] = last
	}
	//fmt.Println(fillArr)

	// 操作原数组
	left, mid, right := 0, 0, arrLen
	// 斐波那契数组的游标
	k := len(fbArr) - 1
	for left <= right {
		fmt.Println("[DEBUG left, right]:\t", left, right)
		//fmt.Println("[DEBUG fbArr[k]]:\t", fbArr[k])
		mid = left + fbArr[k-1] - 1
		if val < fillArr[mid] {
			// 在左边，取 f(k-1)
			right = mid - 1
			k -= 1
		} else if val > fillArr[mid] {
			// 在右边，取 f(k-2)
			left = mid + 1
			k -= 2
		} else {
			if mid > right {
				// 超出原数组，取原数组最后一个位置
				return right
			} else {
				// 值相等，查找成功
				return mid
			}
		}
	}
	// 未查找到
	return -1
}

//
// 生成待查找数组对应的斐波那契数组
//
func makeFibonacciArray(arr []int) []int {
	n := len(arr)
	fbLen := 2
	first, second, third := 1, 2, 3

	// 计算最接近 arr 长度的斐波那契数组的长度
	for third < n {
		// 累加值往后挪
		third, first, second = first+second, second, third
		fbLen++
	}

	// 生成指定长度的斐波那契数组
	fb := make([]int, fbLen)
	fb[0] = 1
	fb[1] = 1
	for i := 2; i < fbLen; i++ {
		fb[i] = fb[i-1] + fb[i-2]
	}
	return fb
}
