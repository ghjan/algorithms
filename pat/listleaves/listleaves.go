package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ghjan/algorithms/binarytree"
	"github.com/ghjan/algorithms/set"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func test1() {
	GOPATH := os.Getenv("GOPATH")
	f := "listleaves_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var n int
	//binaryTree := make([]Node, 0)
	var simpleBinaryTree binarytree.SimpleBinaryTree
	var rootNode int
	var notRootSet set.ItemSet
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 {
			n, _ = strconv.Atoi(string(a))
			if n == 1 {
				rootNode = 0
			} else if n < 1 {
				return
			}
			simpleBinaryTree = make([]binarytree.SimpleNode, n)
		} else //读取节点数据
		{
			sons := strings.Split(string(a), " ")
			left := -1
			right := -1
			if sons[0] != "-" {
				left, _ = strconv.Atoi(sons[0])
			}
			if sons[1] != "-" {
				right, _ = strconv.Atoi(sons[1])
			}
			simpleBinaryTree[i-1] = binarytree.SimpleNode{rune(i - 1), left, right}
			if i == n { //最后一个节点数据
				break
			}
			//if n > 1 && simpleBinaryTree[i-1].IsLeaf() {
			//	notRootSet.Add(i - 1)
			//}
			if left >= 0 {
				notRootSet.Add(left)
			}
			if right >= 0 {
				notRootSet.Add(right)
			}
		}

	}
	result := ""

	//获取root节点
	for i := 0; i < n; i++ {
		if !notRootSet.Has(i) {
			rootNode = i
			break
		}
	}
	simpleBinaryTree.LevelOrderTraverse(rootNode, func(node binarytree.SimpleNode) {
		if node.IsLeaf() {
			result += fmt.Sprintf("%d ", node.Data)
		}
	})
	fmt.Printf(strings.Trim(result, " "))
}

func main() {
	test1()
}
