package binarytree

import (
	"fmt"
	"io"
	"testing"

	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/ghjan/algorithms/set"
	"github.com/stretchr/testify/assert"
)

func initTree() BinaryTree {

	arr := []int{10, 5, 24, 30, 60, 40, 45, 15, 27, 49, 23, 42, 56, 12, 8, 55, 2, 9}
	fmt.Println(arr)
	tree := CreateBinaryTree(arr)

	return tree
}

func TestTraverse(t *testing.T) {
	tree := initTree()
	fmt.Println("----PreOrderTraverse-----------")
	result := ""
	tree.PreOrderTraverse(func(nodeMe *Node) {
		result += fmt.Sprintf("%d ", nodeMe.Data)
	})
	fmt.Println(result)
	assert.Equal(t, "10 5 30 15 55 2 27 9 60 49 23 24 40 42 56 45 12 8 ", result, "")
	fmt.Println("\n----InOrderTraverse-----------")
	result = ""
	tree.InOrderTraverse(func(nodeMe *Node) {
		result += fmt.Sprintf("%d ", nodeMe.Data)
	})
	assert.Equal(t, "55 15 2 30 9 27 5 49 60 23 10 42 40 56 24 12 45 8 ", result, "")
	fmt.Println(result)
	fmt.Println("\n----PostOrderTraverse-----------")
	result = ""
	tree.PostOrderTraverse(func(nodeMe *Node) {
		result += fmt.Sprintf("%d ", nodeMe.Data)
	})
	assert.Equal(t, "55 2 15 9 27 30 49 23 60 5 42 56 40 12 8 45 24 10 ", result, "")

	fmt.Println(result)

	fmt.Println("\n----LevelOrderTraverse-----------")
	result = ""
	tree.LevelOrderTraverse(func(node *Node) {
		result += fmt.Sprintf("%d ", node.Data)
	})
	assert.Equal(t, "10 5 24 30 60 40 45 15 27 49 23 42 56 12 8 55 2 9 ", result, "")
	fmt.Println(result)
}

func TestLevelOrderTraverseSimple(t *testing.T) {
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
	var simpleBinaryTree SimpleBinaryTree
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
			simpleBinaryTree = make([]SimpleNode, n)
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
			simpleBinaryTree[i-1] = SimpleNode{rune(i - 1), left, right}
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
	simpleBinaryTree.LevelOrderTraverse(rootNode, func(node SimpleNode) {
		if node.IsLeaf() {
			result += fmt.Sprintf("%d ", node.Data)
		}
	})
	result = strings.Trim(result, " ")
	fmt.Printf(result)
	assert.Equal(t, "4 1 5", result)
}

func TestLevelOrderTraverse(t *testing.T) {
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
	var tree BinaryTree
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
			tree = make([]Node, n)
		} else //读取节点数据
		{
			sons := strings.Split(string(a), " ")
			left := -1
			right := -1
			var leftChild *Node = nil
			var rightChild *Node = nil
			if sons[0] != "-" {
				left, _ = strconv.Atoi(sons[0])
				leftChild = &tree[left]
			}
			if sons[1] != "-" {
				right, _ = strconv.Atoi(sons[1])
				rightChild = &tree[right]
			}
			if left >= 0 {

			}

			tree[i-1] = Node{rune(i - 1), leftChild, rightChild}
			if i == n { //最后一个节点数据
				break
			}
			//if n > 1 && tree[i-1].IsLeaf() {
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
	tree[rootNode].LevelOrderTraverse(func(node *Node) {
		if node.IsLeaf() {
			result += fmt.Sprintf("%d ", node.Data)
		}
	})
	result = strings.Trim(result, " ")
	fmt.Printf(result)
	assert.Equal(t, "4 1 5", result)
}
