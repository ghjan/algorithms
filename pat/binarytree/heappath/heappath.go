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
/*
https://pintia.cn/problem-sets/951072707007700992/problems/980373149213519872

05-树7 堆中的路径（25 分）

将一系列给定数字插入一个初始为空的小顶堆H[]。随后对任意给定的下标i，打印从H[i]到根结点的路径。

输入格式:

每组测试第1行包含2个正整数N和M(≤1000)，分别是插入元素的个数、以及需要打印的路径条数。下一行给出区间[-10000, 10000]内的N个要被插入一个初始为空的小顶堆的整数。最后一行给出M个下标。

输出格式:

对输入中给出的每个下标i，在一行中输出从H[i]到根结点的路径上的数据。数字间以1个空格分隔，行末不得有多余空格。
 */
func solveHeapPath(filename string) {

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
	solveHeapPath(filename)
}
