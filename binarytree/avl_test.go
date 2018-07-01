package binarytree

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestAvl(t *testing.T) {
	data := []int{3, 2, 1, 4, 5, 6, 7, 16, 15, 14, 13, 12, 11, 10, 8}
	fmt.Println(data)
	root := NewAVLTreeNode(9)
	var err error
	for _, value := range data {
		if root, err = root.Insert(value); err != nil {
			fmt.Println(err)
		} else {
			fmt.Print(" root -> Key:", root.Key, ", height:", root.height)
			if root.left != nil {
				fmt.Print(", left:", root.left.Key)
			}
			if root.right != nil {
				fmt.Print(", right:", root.right.Key)
			}
			fmt.Println()
		}
	}

	// fmt.Println(root.DisplayAsc())
	// fmt.Println(root.DisplayDesc())

	root.ProcessDepth(0)
	depth := -1
	root.LevelOrderTraverse(func(avlNode *AvlTreeNode) {
		if avlNode.depth == depth+1 {
			fmt.Println()
			depth++
		}
		fmt.Printf("%d \t", avlNode.Key)
	})
}

func TestAvlRoot(t *testing.T) {
	filename := strings.Join([]string{"E:/go-work/bin", "avl_case_1.txt"}, "/")

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var n int
	var root *AvlTreeNode
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { // n is the total number of keys to be inserted.
			n, _ = strconv.Atoi(string(a))
			//if n == 1 {
			//	rootNode = 0
			//} else if n < 1 {
			//	return
			//}
			//simpleBinaryTree = make([]SimpleNode, n)
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
				} else {
				fmt.Print(" root -> Key:", root.Key, ", height:", root.height)
				if root.left != nil {
				fmt.Print(", left:", root.left.Key)
				}
				if root.right != nil {
				fmt.Print(", right:", root.right.Key)
				}
				fmt.Println()
				}
			}
		}
	}
	fmt.Print(root.Key)
}
