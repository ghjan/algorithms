package inthashtable

import (
	"strings"
	"strconv"
	"fmt"
	"os"
	"bufio"
	"io"
)

func PrintPosition(positions []int) {
	for index, pos := range positions {
		if pos < 0 {
			fmt.Printf("%s", "-")
		} else {
			fmt.Printf("%d", pos)
		}
		if index < len(positions)-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}

}

//CreateTableForHashing 为hashing创建hash table
//isLinear 是否开放地址法的线性探测
func CreateTableForHashing(filename string, isLinear bool) []int {
	var result []int
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()

	var M, N int
	var table IntHashTable
	br := bufio.NewReader(fi)
	i := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //user defined table size， the number of input numbers
			array := strings.Split(string(a), " ")
			M, _ = strconv.Atoi(array[0])
			N, _ = strconv.Atoi(array[1])
			table = CreateHashTable(M)

		} else { // hash table 输入数据
			array := strings.Split(string(a), " ")
			for index, item := range array {
				if index >= N {
					break
				}
				intValue, _ := strconv.Atoi(item)
				pos := -1
				if isLinear {
					pos = table.InsertLinear(intValue)
				} else {
					pos = table.Insert(intValue)
				}
				result = append(result, pos)
			}

		}
		i++
	}
	return result
}

func CreateTableForHashingHard(filename string, isLinear bool) IntHashTable {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()

	var N int
	var table IntHashTable
	br := bufio.NewReader(fi)

	for i := 0; i < 2; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //hash table大小
			N, _ = strconv.Atoi(string(a))
			table = CreateHashTable(N - 1)

		} else { //hash table的已有状态
			array := strings.Split(string(a), " ")
			for index, item := range array {
				if index >= N {
					break
				}
				intValue, _ := strconv.Atoi(item)
				table.Cells[index].Data = intValue
				if intValue >= 0 {
					table.Cells[index].Info = Legitimate
				} else { //负数表示没有使用这个位置
					table.Cells[index].Info = Empty
				}
			}
		}
	}

	return table
}
