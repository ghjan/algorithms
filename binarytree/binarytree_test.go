package binarytree

import (
"fmt"
"testing"

"github.com/stretchr/testify/assert"
)

func initTree() BinaryTree {

	arr := []int{10, 5, 24, 30, 60, 40, 45, 15, 27, 49, 23, 42, 56, 12, 8, 55, 2, 9}
	fmt.Println(arr)
	tree := CreateTree(arr)

	return tree
}

func TestTraverse(t *testing.T) {
	tree := initTree()
	fmt.Println("----PreOrderTraverse-----------")
	result := ""
	PreOrderTraverse(&tree[0], func(nodeMe *Node) {
		result += fmt.Sprintf("%d ", nodeMe.data)
	})
	fmt.Println(result)
	assert.Equal(t, "10 5 30 15 55 2 27 9 60 49 23 24 40 42 56 45 12 8 ", result, "")
	fmt.Println("\n----InOrderTraverse-----------")
	result = ""
	InOrderTraverse(&tree[0], func(nodeMe *Node) {
		result += fmt.Sprintf("%d ", nodeMe.data)
	})
	assert.Equal(t, "55 15 2 30 9 27 5 49 60 23 10 42 40 56 24 12 45 8 ", result, "")
	fmt.Println(result)
	fmt.Println("\n----PostOrderTraverse-----------")
	result = ""
	PostOrderTraverse(&tree[0], func(nodeMe *Node) {
		result += fmt.Sprintf("%d ", nodeMe.data)
	})
	assert.Equal(t, "55 2 15 9 27 30 49 23 60 5 42 56 40 12 8 45 24 10 ", result, "")

	fmt.Println(result)

	fmt.Println("\n----LevelOrderTraverse-----------")
	result = ""
	LevelOrderTraverse(&tree[0], func(node *Node) {
		result += fmt.Sprintf("%d ", node.data)
	})
	assert.Equal(t, "10 5 24 30 60 40 45 15 27 49 23 42 56 12 8 55 2 9 ", result, "")
	fmt.Println(result)
}
