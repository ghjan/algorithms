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
/*
AVL:自平衡二叉查找树
Root of AVL Tree  2013年浙江大学计算机学院免试研究生上机考试真题，是关于AVL树的基本训练，一定要做；
http://pintia.cn/problem-sets/951072707007700992/problems/977489194356715520
 */
func solveAVL(filename string) {

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
	GOPATH := os.Getenv("GOPATH")
	for _, f := range fileNames {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		solveAVL(filename)
		fmt.Println()
	}
}
