package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ghjan/algorithms/binarytree"
	"github.com/ghjan/algorithms/set"
)

func buildList(filename string) (binarytree.SimpleBinaryTree, int) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, -1
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var n int
	//binaryTree := make([]Node, 0)
	var simpleBinaryTree binarytree.SimpleBinaryTree
	var rootNode int
	var notRootSet set.ItemSet //非根节点的集合
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
				return nil, -1
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
	//获取root节点
	for i := 0; i < n; i++ {
		if !notRootSet.Has(i) {
			rootNode = i
			break
		}
	}

	return simpleBinaryTree, rootNode
}
func solveListLeaves() error {
	GOPATH := os.Getenv("GOPATH")
	f := "listleaves_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	simpleBinaryTree, rootNode := buildList(filename)
	if simpleBinaryTree == nil || rootNode < 0 {
		return errors.New("simpleBinaryTree is nil or rootNode is -1")
	}
	result := ""

	simpleBinaryTree.LevelOrderTraverse(rootNode, func(node binarytree.SimpleNode) {
		if node.IsLeaf() {
			result += fmt.Sprintf("%d ", node.Data)
		}
	})
	fmt.Printf(strings.Trim(result, " "))
	return nil
}

func main() {
	if err := solveListLeaves(); err != nil {
		fmt.Println(err)
	}
}
