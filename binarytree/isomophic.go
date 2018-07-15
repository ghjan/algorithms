package binarytree

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

func buildTrees(filename string) (BinaryTree, BinaryTree) {

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, nil
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var N int
	var tree, tree2 BinaryTree
	begin := true
	index := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if begin { // n is the total number of keys to be inserted.
			if N, err = strconv.Atoi(string(a)); N > 0 && err == nil {
				if index == 0 {
					tree = InitNewTree(N)
				}
				//else {
				//	tree2 = InitNewTree(N)
				//}
				begin = false
			} else {
				fmt.Printf("N:%d", N)
				fmt.Println(err)
				break
			}
		} else {
			if nodeInfo := strings.Split(string(a), " "); len(nodeInfo) >= 2 {
				leftIndex := -1
				rightIndex := -1
				if nodeInfo[1] != "-" {
					leftIndex, _ = strconv.Atoi(nodeInfo[1])
				}
				if nodeInfo[2] != "-" {
					rightIndex, _ = strconv.Atoi(nodeInfo[2])
				}
				if index <= N {
					tree.Insert(rune(nodeInfo[0][0]), (index-1)%N, leftIndex, rightIndex)
				} else {
					tree2.Insert(rune(nodeInfo[0][0]), (index-2)%N, leftIndex, rightIndex)
				}
			} else {
				tree2 = InitNewTree(N)
			}
		}
		index++
	}
	return tree, tree2
}
func SolveBTIsomorphic(filename string) {
	tree, tree2 := buildTrees(filename)
	root1 := tree.RootIndex()
	root2 := tree2.RootIndex()
	if tree.Isomorphic(tree2, root1, root2) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
