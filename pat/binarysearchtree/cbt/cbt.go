package main

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"bufio"
	"os"
	"strconv"

	"github.com/ghjan/algorithms/binarysearchtree/cbt"
)

/*
cbt:Complete Binary Search Tree

Complete Binary Search Tree 2013年秋季PAT甲级真题，略有难度，量力而行。第7周将给出讲解。
http://pintia.cn/problem-sets/951072707007700992/problems/977489256881188864
04-树6 Complete Binary Search Tree（30 分）

A Binary Search Tree (BST) is recursively defined as a binary tree which has the following properties:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.
A Complete Binary Tree (CBT) is a tree that is completely filled, with the possible exception of the bottom level,
which is filled from left to right.

Now given a sequence of distinct non-negative integer keys, a unique BST can be constructed if it is required
that the tree must also be a CBT. You are supposed to output the level order traversal sequence of this BST.
*/
func initData(fileName string) []int {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var N int
	var result []int
	for i := 0; i < 2; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { // n is the total number
			N, _ = strconv.Atoi(string(a))
			result = make([]int, N)
		} else {
			array1 := strings.Split(string(a), " ")
			for i := 0; i < len(array1); i++ {
				result[i], _ = strconv.Atoi(array1[i])
			}
		}
	}
	return result
}

func solveAndLevelTraversal(fileName string) {
	A := initData(fileName)
	sort.Ints(A)
	N := len(A)
	T := make([]int, N)
	cbt.SolveCBT(A, T, 0, N-1, 0)

	//顺序打印就可以输出层级遍历
	// level order traversal sequence of this BST.
	result := ""
	for i := 0; i < N; i++ {
		result += fmt.Sprintf("%d ", T[i])
	}
	fmt.Println(strings.TrimRight(result, " "))
}

func main() {
	GOPATH := os.Getenv("GOPATH")
	f := "completebst_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	solveAndLevelTraversal(filename)

}
