package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/ghjan/algorithms/binarytree"
)

func test1(filename string) {

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var n int
	var root *binarytree.AvlTreeNode
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { // n is the total number of keys to be inserted.
			n, _ = strconv.Atoi(string(a))
		} else //读取节点数据
		{
			data := strings.Split(string(a), " ")
			for index, value := range data {
				if index >= n {
					break
				}
				intValue, _ := strconv.Atoi(value)
				if root, err = root.Insert(intValue); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	fmt.Print(root.Key)
}

func main() {
	fileNames := [...]string{"avl_case_1.txt", "avl_case_2.txt"}
	for _, f := range fileNames {
		filename := strings.Join([]string{"E:/go-work/bin", f}, "/")
		test1(filename)
		fmt.Println()
	}
}
