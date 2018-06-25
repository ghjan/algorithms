package binarysearchtree

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
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
	tree.Insert(14, "14")
	tree.Insert(12, "12")
	tree.Insert(13, "13")
	tree.Insert(11, "11")
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
	if fmt.Sprintf("%s", max) != "14" {
		t.Errorf("Max() should return 14 but return %s", max)
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

	fmt.Println("----删除一个不存在的-----")
	//删除一个不存在的
	node_removed, newNode := treeLocal.Remove(222)
	treeLocal.String()
	if node_removed != nil {
		t.Errorf("remove can not return expected node_removed:%v, actual:%v", nil, node_removed)
	}
	if newNode != nil {
		t.Errorf("remove can not return expected newNode:%v, actual:%v", nil, newNode)
	}

	fmt.Println("----删除最左边的-----")
	//删除最左边的
	node_removed, newNode = treeLocal.Remove(1)
	treeLocal.String()
	assert.Equal(t, node_removed.key, 1, fmt.Sprintf("remove can not return expected node_removed which has value:%d, actual:%d", 1, node_removed.key))
	if newNode != nil {
		t.Errorf("remove can not return expected newNode:%v, actual:%v", nil, newNode)
	}
	min_value := fmt.Sprintf("%s", *treeLocal.Min())
	if min_value != "2" {
		t.Errorf("Remove(1) failed,expected min_value:%s, actual:%s", "2", min_value)
	}

	fmt.Println("----删除有两个儿子的结点-----")
	//删除有两个儿子的结点
	node_removed, newNode = treeLocal.Remove(6)
	treeLocal.String()
	assert.Equal(t, newNode.key, 7, fmt.Sprintf("remove can not return expected newNode:%d, actual:%d", 7, newNode.key))
	assert.Equal(t, node_removed.key, 6, fmt.Sprintf("remove can not return expected node_removed which has value:%d, actual:%d", 6, node_removed.key))

	fmt.Println("----删除最右边的-----")
	//删除最右边的
	max_value_s := fmt.Sprintf("%s", *treeLocal.Max())

	if max_value, error := strconv.Atoi(max_value_s); error != nil {
		t.Errorf("max_value_s:%v", max_value_s)
	} else {
		node_removed, newNode = treeLocal.Remove(max_value)
		treeLocal.String()
		assert.Equal(t, newNode.key, 12, fmt.Sprintf("remove can not return expected newNode:%d, actual:%d", 12, newNode.key))
		assert.Equal(t, node_removed.key, max_value, fmt.Sprintf("remove can not return expected node_removed which has value:%d, actual:%d", max_value, node_removed.key))
		max_value_s = fmt.Sprintf("%s", *treeLocal.Max())
		if max_value2, error := strconv.Atoi(max_value_s); error != nil {
			t.Errorf("max_value_s:%v", max_value_s)
		} else {
			if max_value2 != max_value-1 {
				t.Errorf("Remove(1) failed,expected max_value2:%s, actual:%s", max_value-1, max_value2)
			}
		}
	}

}
