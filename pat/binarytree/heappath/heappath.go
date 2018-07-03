package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"bufio"
	"os"

	"github.com/ghjan/algorithms/heap"
)

func TestHeap_Path(filename string) {

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var n, m int //插入元素的个数、以及需要打印的路径条数
	var keysString, positionsString string
	for i := 0; i < 3; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		switch i {
		case 0:
			nm := strings.Split(string(a), " ")
			n, _ = strconv.Atoi(nm[0])
			m, _ = strconv.Atoi(nm[1])
			break
		case 1: // 给出区间[-10000, 10000]内的N个要被插入一个初始为空的小顶堆的整数
			keysString = string(a)
			break
		case 2: //最后一行给出M个下标
			positionsString = string(a)
			break
		}
	}
	keys := strings.Split(keysString, " ")
	H := heap.Create(len(keys))
	for index, key := range keys {
		if index >= n {
			break
		}
		k, _ := strconv.Atoi(key)
		H.Insert(k)
	}

	positions := strings.Split(positionsString, " ")
	for index, search := range positions {
		if index >= m {
			break
		}
		X, _ := strconv.Atoi(search)
		fmt.Println(H.Path(X))
	}
}

func main() {
	GOPATH := os.Getenv("GOPATH")
	f := "heappath_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	TestHeap_Path(filename)
}
