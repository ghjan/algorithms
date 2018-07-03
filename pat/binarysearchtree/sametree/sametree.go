package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/ghjan/algorithms/binarysearchtree"
)

func test1(filename string) {

	var treeLocal binarysearchtree.ItemBinarySearchTree
	var treeGenerated binarysearchtree.ItemBinarySearchTree

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var N, L int
	begin := true
	index := 0
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if begin { // n is the total number of keys to be inserted.
			numbers := strings.Split(string(a), " ")
			if len(numbers) >= 2 {
				N, _ = strconv.Atoi(string(numbers[0]))
				L, _ = strconv.Atoi(string(numbers[1]))
				begin = false
				index = 0
			} else {
				break
			}
		} else {
			array1 := strings.Split(string(a), " ")
			if index == 0 { // 初始插入序列
				treeLocal.FactoryFromArray2(array1[:N])
			} else if index > 0 && index <= L { //L个需要检查的序列
				treeGenerated.FactoryFromArray2(array1[:N])
				if treeLocal.Equal(&treeGenerated) {
					fmt.Println("YES")
				} else {
					fmt.Println("NO")
				}
				treeGenerated.Destroy()
				if index == L { //已经是最后一个需要检查的序列
					treeLocal.Destroy()
					begin = true
				}
			}
			index++
		}
	}
}

func main() {
	GOPATH := os.Getenv("GOPATH")
	f := "bst_sametree_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	test1(filename)

}
