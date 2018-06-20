package main

import (
	"algorithms"
	"fmt"
)

func main() {
	arr := algorithms.GetRange(1, 10000)
	fmt.Println(sequenceSearch(arr, 233))
}

func sequenceSearch(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
