package binarysearchtree

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

var tree ItemBinarySearchTree

func initTree(tree *ItemBinarySearchTree) {
	tree.Insert(8, "8")
	tree.Insert(4, "4")
	tree.Insert(10, "10")
	tree.Insert(2, "2")
	tree.Insert(6, "6")
	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(5, "5")
	tree.Insert(7, "7")
	tree.Insert(9, "9")
}

func TestInsert(t *testing.T) {
	initTree(&tree)
	tree.String()
	tree.Insert(11, "11")
	tree.String()
}

func TestPreOrderTraverse(t *testing.T) {
	traverse := ""
	tree.PreOrderTraverse(func(value Item) {
		traverse += fmt.Sprintf("%s\t", value)
	})
	println(traverse)
}

func TestInOrderTraverse(t *testing.T) {
	traverse := ""
	tree.InOrderTraverse(func(value Item) {
		traverse += fmt.Sprintf("%s\t", value)
	})
	println(traverse)
}

func TestPostOrderTraverse(t *testing.T) {
	traverse := ""
	tree.PostOrderTraverse(func(value Item) {
		traverse += fmt.Sprintf("%s\t", value)
	})
	println(traverse)
}

func TestMin(t *testing.T) {
	min := *tree.Min()
	if fmt.Sprintf("%s", min) != "1" {
		t.Errorf("Min() should return 1 but return %s", min)
	}
}

func TestMax(t *testing.T) {
	max := *tree.Max()
	if fmt.Sprintf("%s", max) != "11" {
		t.Errorf("Max() should return 11 but return %s", max)
	}
}

func TestSearch(t *testing.T) {
	for i := 1; i <= 11; i++ {
		if !tree.Search(i) {
			t.Errorf("Search() can't find %d", i)
		}
	}
}

func TestRemove(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	initTree(&treeLocal)
	treeLocal.String()

	//删除一个不存在的
	node_removed, newNode := treeLocal.Remove(11)
	treeLocal.String()
	assert.Equal(t, node_removed, nil, fmt.Sprintf("remove can not return expected node_removed:%v, actual:%v", nil, node_removed))
	assert.Equal(t, newNode, nil, fmt.Sprintf("remove can not return expected newNode:%v, actual:%v", nil, newNode))

	//删除最左边的
	node_removed, newNode = treeLocal.Remove(1)
	treeLocal.String()
	assert.Equal(t, node_removed.key, 1, fmt.Sprintf("remove can not return expected node_removed which has value:%d, actual:%d", 1, node_removed.key))
	assert.Equal(t, newNode, nil, fmt.Sprintf("remove can not return expected newNode:%v, actual:%v", nil, newNode))
	if fmt.Sprintf("%s", *treeLocal.Min()) != "2" {
		t.Errorf("Remove(1) failed")
	}

	//删除有两个儿子的结点
	node_removed, newNode = treeLocal.Remove(6)
	treeLocal.String()
	assert.Equal(t, node_removed.key, 6, fmt.Sprintf("remove can not return expected node_removed which has value:%d, actual:%d", 6, node_removed.key))
	assert.Equal(t, newNode.key, 7, fmt.Sprintf("remove can not return expected newNode:%d, actual:%d", 7, newNode.key))

}
