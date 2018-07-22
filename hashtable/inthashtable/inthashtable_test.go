package inthashtable

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func printPosition(positions []int) {
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
func TestSolveHashing(t *testing.T) {
	//N个记录
	var M, N int
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"hashing_case_1.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		positionSlice := createHashtableFromFile(filename, N, M, false)
		printPosition(positionSlice)
	}
}

func createHashtableFromFile(filename string, N int, M int, isLinear bool) []int {
	var result []int
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()
	var table IntHashTable
	br := bufio.NewReader(fi)
	i := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //顶点数量， 边数量
			array := strings.Split(string(a), " ")
			N, _ = strconv.Atoi(array[0])
			M, _ = strconv.Atoi(array[1])
			table = CreateHashTable(N)

		} else { //边的数据输入 start, end, weight
			//单向的
			array := strings.Split(string(a), " ")
			for index, item := range array {
				if index >= M {
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
