package main

import (
	"fmt"
)

func main() {
	mapResults := make(map[int]string)
	var arrResults [][]string
	count := 5
	for i := 0; i < count; i++ {
		valueStr := fmt.Sprintf("this is %d", i)
		mapResults[i] = valueStr
		var tmpArr []string
		for j := 0; j < 15; j++ {
			tmpArr = append(tmpArr, "a")
		}
		arrResults = append(arrResults, tmpArr)
	}
	fmt.Println(mapResults)
	fmt.Println(arrResults)

	arrResults[4][10] = "b"
	fmt.Println(arrResults)
}
