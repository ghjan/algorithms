package binarysearchtree

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"testing"

	"bufio"
	"os"
	"strings"

	"github.com/stretchr/testify/assert"
)

func initTree(tree *ItemBinarySearchTree) {
	array1 := []int{8, 4, 10, 2, 6, 1, 3, 5, 7, 9, 14, 12, 13, 11}
	tree.FactoryFromArray(array1)
}

func TestFactory(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	fmt.Println("---3, 1, 4, 2---------")
	array1 := []int{3, 1, 4, 2}
	treeLocal.FactoryFromArray(array1)
	treeLocal.LevelOrderTraverse(func(item Item) {
		fmt.Print(item)
	})
	treeLocal.Destroy()
	fmt.Println()
	fmt.Println("-----2, 1-------")
	array2 := []int{2, 1}
	treeLocal.FactoryFromArray(array2)
	treeLocal.LevelOrderTraverse(func(item Item) {
		fmt.Print(item)
	})

}

func TestInsert(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	initTree(&treeLocal)
	treeLocal.String()

	fmt.Println("------TestInsert after init--------")
	treeLocal.String()
	if result := treeLocal.Insert(11, "11"); result != nil {
		t.Errorf("Insert(11) failed,expected return:%v, actual:%v", nil, result)
	}
	fmt.Println("------TestInsert after Insert(11)--------")
	treeLocal.String()
	if result := treeLocal.Insert(20, "20"); result == nil || result.key != 14 {
		t.Errorf("Insert(20) failed,expected return:%d, actual:%d", 14, result.key)
	}
	fmt.Println("------TestInsert after Insert(20)--------")
	treeLocal.String()
}

func TestPreOrderTraverse(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	initTree(&treeLocal)

	traverse := ""
	treeLocal.PreOrderTraverse(func(value Item) {
		traverse += fmt.Sprintf("%s\t", value)
	})
	fmt.Println(traverse)
}

func TestInOrderTraverse(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	initTree(&treeLocal)
	traverse := ""
	treeLocal.InOrderTraverse(func(value Item) {
		traverse += fmt.Sprintf("%s\t", value)
	})
	fmt.Println(traverse)
}

func TestPostOrderTraverse(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	initTree(&treeLocal)
	traverse := ""
	treeLocal.PostOrderTraverse(func(value Item) {
		traverse += fmt.Sprintf("%s\t", value)
	})
	fmt.Println(traverse)
}

func TestLevelOrderTraverse(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	initTree(&treeLocal)
	traverse := ""
	treeLocal.LevelOrderTraverse(func(value Item) {
		traverse += fmt.Sprintf("%s\t", value)
	})
	fmt.Println(traverse)
}

func TestMin(t *testing.T) {
	var treeLocal ItemBinarySearchTree
	initTree(&treeLocal)

	min := treeLocal.Min()
	if min == nil || fmt.Sprintf("%s", *min) != "1" {
		t.Errorf("Min() should return 1 but return %v", min)
	}
}

func TestMax(t *testing.T) {
	var treeLocal ItemBinarySearchTree
	initTree(&treeLocal)

	max := *treeLocal.Max()
	if fmt.Sprintf("%s", max) != "14" {
		t.Errorf("Max() should return 14 but return %s", max)
	}
}

func TestSearch(t *testing.T) {
	var treeLocal ItemBinarySearchTree
	initTree(&treeLocal)
	for i := 1; i <= 14; i++ {
		if node, success := treeLocal.Search(i); !success {
			t.Errorf("Search() can't find %d", i)
		}else{
			//fmt.Printf("node found:%v\n", node)
			assert.Equal(t, i, node.key)
			assert.Equal(t, strconv.Itoa(i), node.value)
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
	maxValue_s := fmt.Sprintf("%s", *treeLocal.Max())

	if maxValue, error := strconv.Atoi(maxValue_s); error != nil {
		t.Errorf("maxValue_s:%v", maxValue_s)
	} else {
		node_removed, newNode = treeLocal.Remove(maxValue)
		treeLocal.String()
		assert.Equal(t, newNode.key, 12, fmt.Sprintf("remove can not return expected newNode:%d, actual:%d", 12, newNode.key))
		assert.Equal(t, node_removed.key, maxValue, fmt.Sprintf("remove can not return expected node_removed which has value:%d, actual:%d", maxValue, node_removed.key))
		maxValue_s = fmt.Sprintf("%s", *treeLocal.Max())
		if maxValue2, error := strconv.Atoi(maxValue_s); error != nil {
			t.Errorf("maxValue_s:%v", maxValue_s)
		} else {
			if maxValue2 != maxValue-1 {
				t.Errorf("Remove(1) failed,expected maxValue2:%d, actual:%d", maxValue-1, maxValue2)
			}
		}
	}

}

func TestEqual(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	initTree(&treeLocal)

	var treeGenerated ItemBinarySearchTree

	treeLocal.PreOrderTraverse(func(value Item) {
		if v, err := strconv.Atoi(fmt.Sprintf("%s", value)); err != nil {
			t.Errorf("value:%s", value)
		} else {
			treeGenerated.Insert(v, value)
		}
	})
	assert.Equal(t, treeLocal.Equal(&treeGenerated), true, fmt.Sprintf("two trees are expected to be Equal , but actually they are not Equal"))
}

func TestIsomorphic(t *testing.T) {
	var treeLocal ItemBinarySearchTree

	initTree(&treeLocal)

	var treeGenerated ItemBinarySearchTree

	treeLocal.PreOrderTraverse(func(value Item) {
		if v, err := strconv.Atoi(fmt.Sprintf("%s", value)); err != nil {
			t.Errorf("value:%s", value)
		} else {
			treeGenerated.Insert(v, value)
		}
	})
	fmt.Println("----------TestIsomorphic 1st part-----------")
	treeLocal.String()
	fmt.Println("-------treeGenerated---------")
	treeGenerated.String()
	assert.Equal(t, treeLocal.Isomorphic(&treeGenerated), true, fmt.Sprintf("two trees are expected to be isomorphic , but actually they are not isomorphic"))

	treeGenerated.PreOrderTraverse2(func(nodeMe *Node) {
		if nodeMe != nil && rand.Intn(20)/2 == 0 { //随机交换左右子树
			nodeMe.left, nodeMe.right = nodeMe.right, nodeMe.left
		}
	})
	fmt.Println("----------TestIsomorphic 2nd part-----------")
	treeLocal.String()
	fmt.Println("-------treeGenerated---------")
	treeGenerated.String()
	assert.Equal(t, treeLocal.Isomorphic(&treeGenerated), true, fmt.Sprintf("two trees are expected to be isomorphic , but actually they are not isomorphic"))

}

func TestSameTree(t *testing.T) {
	f := "bst_sametree_case_1.txt"
	GOPATH := os.Getenv("GOPATH")
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")

	var treeLocal ItemBinarySearchTree
	var treeGenerated ItemBinarySearchTree

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var N, L int
	begin := true
	index := 0
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if begin { // n is the total number of keys to be inserted.
			numbers := strings.Split(string(a), " ")
			if len(numbers) >= 2 {
				N, _ = strconv.Atoi(string(numbers[0]))
				L, _ = strconv.Atoi(string(numbers[1]))
				begin = false
				index = 0
			} else {
				break
			}
		} else {
			array1 := strings.Split(string(a), " ")
			if index == 0 { // 初始插入序列
				treeLocal.FactoryFromArray2(array1[:N])
			} else if index > 0 && index <= L { //L个需要检查的序列
				treeGenerated.FactoryFromArray2(array1[:N])
				if treeLocal.Equal(&treeGenerated) {
					fmt.Println("YES")
				} else {
					fmt.Println("NO")
				}
				treeGenerated.Destroy()
				if index == L { //已经是最后一个需要检查的序列
					treeLocal.Destroy()
					begin = true
				}
			}
			index++
		}
	}
}
