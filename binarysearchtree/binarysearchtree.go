package binarysearchtree

import (
	"fmt"
	"sync"
	"github.com/cheekybits/genny/generic"
	"strconv"
)

type Item generic.Type

type Node struct {
	key   int   // 先序遍历的节点序号
	value Item  // 节点存储的值
	left  *Node // 左子节点
	right *Node // 右子节点
}

type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (node *Node) Equal(target *Node) bool {
	if node == nil || target == nil {
		return node == nil && target == nil
	} else {
		return node.key == target.key && node.value == target.value && node.left.Equal(target.left) && node.right.Equal(target.right)
	}
}

func (node *Node) SelfEqual(target *Node) bool {
	if node == nil || target == nil {
		return node == nil && target == nil
	} else {
		return node.key == target.key && node.value == target.value
	}
}

func (tree *ItemBinarySearchTree) FactoryFromArray(arr []int) {
	for _, item := range arr {
		tree.Insert(item, strconv.Itoa(item))
	}
}

func (tree *ItemBinarySearchTree) FactoryFromArray2(arr []string) {
	for _, item := range arr {
		intValue, _ := strconv.Atoi(item)
		tree.Insert(intValue, item)
	}
}

//释放
func (tree *ItemBinarySearchTree) Destroy() {
	if tree == nil {
		return
	}
	if tree != nil {
		tree.PostOrderTraverse2(func(nodeMe *Node) {
			nodeMe = nil
		})
	}
	tree.root = nil
}

// 向树中插入元素
// 返回插入点，即新节点的父亲节点
func (tree *ItemBinarySearchTree) Insert(key int, value Item) *Node {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	newNode := &Node{key, value, nil, nil}
	// 初始化树
	if tree.root == nil {
		tree.root = newNode
		return tree.root
	} else {
		// 在树中递归查找正确的位置并插入
		return insertNode(tree.root, newNode)
	}
}

// 返回插入点，即新节点的父亲节点
func insertNode(node, newNode *Node) *Node {
	// 插入到左子树
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
			return node
		} else {
			// 递归查找左边插入
			return insertNode(node.left, newNode)
		}
	} else if newNode.key > node.key {
		// 插入到右子树
		if node.right == nil {
			node.right = newNode
			return node
		} else {
			// 递归查找右边插入
			return insertNode(node.right, newNode)
		}
	} else {
		return nil
	}
}

// 搜索序号
// 返回true说明找到了
func (tree *ItemBinarySearchTree) Search(key int)  (*Node, bool) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return search(tree.root, key)
}
func search(node *Node, key int) (*Node, bool) {
	if node == nil {
		return nil, false
	}
	// 向左搜索更小的值
	if key < node.key {
		return search(node.left, key)
	}
	// 向右搜索更大的值
	if key > node.key {
		return search(node.right, key)
	}
	return node, true // key == node.key
}

// 删除节点
//返回被删除的结点
func (tree *ItemBinarySearchTree) Remove(key int) (*Node, *Node) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	return remove(nil, tree.root, key, true)
}

// 递归删除节点
func remove(parent, node *Node, key int, isLeft bool) (*Node, *Node) {
	// 要删除的节点不存在
	if node == nil {
		return nil, nil
	}

	// 寻找节点
	// 要删除的节点在左侧
	if key < node.key {
		return remove(node, node.left, key, true)
	}
	// 要删除的节点在右侧
	if key > node.key {
		return remove(node, node.right, key, false)
	}

	// 判断节点类型
	// 要删除的节点是叶子节点，直接删除
	// if key == node.key {
	if node.left == nil && node.right == nil {
		processRemove(parent, nil, isLeft)
		return node, nil
	}

	// 要删除的节点只有一个节点，删除自身
	if node.left == nil {
		processRemove(parent, node.right, isLeft)
		return node, node.right
	}
	if node.right == nil {
		processRemove(parent, node.left, isLeft)
		return node, node.left
	}

	// 要删除的节点有 2 个子节点，找到右子树的最左节点，替换当前节点
	mostLeftNode := node.right
	for {
		// 一直遍历找到最左节点
		if mostLeftNode != nil && mostLeftNode.left != nil {
			mostLeftNode = mostLeftNode.left
		} else {
			break
		}
	}
	// 使用右子树的最左节点替换当前节点，即删除当前节点
	nodeRem := *node
	node.key, node.value = mostLeftNode.key, mostLeftNode.value
	_, mostLeftNode.right = remove(node, node.right, mostLeftNode.key, false) //在当前结点的右边删除mostLeftNode
	return &nodeRem, mostLeftNode
}
func processRemove(parent, node *Node, isLeft bool) {
	if parent != nil {
		if isLeft {
			parent.left = node
		} else {
			parent.right = node
		}
	}
}

// 获取树中值最小的节点：最左节点
func (tree *ItemBinarySearchTree) Min() *Item {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	node := tree.root
	if node == nil {
		return nil
	}
	for {
		if node.left == nil {
			return &node.value
		}
		node = node.left
	}
}

// 获取树中值最大的节点：最右节点
func (tree *ItemBinarySearchTree) Max() *Item {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	node := tree.root
	if node == nil {
		return nil
	}
	for {
		if node.right == nil {
			return &node.value
		}
		node = node.right
	}
}

// 先序遍历(打印函数)：根节点 -> 左子树 -> 右子树
func (tree *ItemBinarySearchTree) PreOrderTraverse(printFunc func(Item)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	preOrderTraverse(tree.root, printFunc)
}
func preOrderTraverse(node *Node, printFunc func(Item)) {
	if node != nil {
		printFunc(node.value)                   // 先打印根结点
		preOrderTraverse(node.left, printFunc)  // 再打印左子树
		preOrderTraverse(node.right, printFunc) // 最后打印右子树
	}
}

// 先序遍历2(操作函數)：根节点 -> 左子树 -> 右子树
func (tree *ItemBinarySearchTree) PreOrderTraverse2(operationFunc func(nodeMe *Node)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	preOrderTraverse2(tree.root, operationFunc)
}
func preOrderTraverse2(node *Node, operationFunc func(nodeMe *Node)) {
	if node != nil {
		operationFunc(node)                          // 先打印根结点
		preOrderTraverse2(node.left, operationFunc)  // 再打印左子树
		preOrderTraverse2(node.right, operationFunc) // 最后打印右子树
	}
}

// 后序遍历（打印函數）：左子树 -> 根节点 -> 右子树
func (tree *ItemBinarySearchTree) PostOrderTraverse(printFunc func(Item)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	postOrderTraverse(tree.root, printFunc)
}
func postOrderTraverse(node *Node, printFunc func(Item)) {
	if node != nil {
		postOrderTraverse(node.left, printFunc)  // 先打印左子树
		postOrderTraverse(node.right, printFunc) // 再打印右子树
		printFunc(node.value)                    // 最后打印根结点
	}
}

// 后序遍历2（操作函數）：左子树 -> 根节点 -> 右子树
func (tree *ItemBinarySearchTree) PostOrderTraverse2(operationFunc func(nodeMe *Node)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	postOrderTraverse2(tree.root, operationFunc)
}
func postOrderTraverse2(node *Node, operationFunc func(nodeMe *Node)) {
	if node != nil {
		postOrderTraverse2(node.left, operationFunc)  // 先打印左子树
		postOrderTraverse2(node.right, operationFunc) // 再打印右子树
		operationFunc(node)                           // 最后打印根结点
	}
}

// 中续遍历：左子树 -> 右子树 -> 根结点
func (tree *ItemBinarySearchTree) InOrderTraverse(printFunc func(Item)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverse(tree.root, printFunc)
}
func inOrderTraverse(node *Node, printFunc func(Item)) {
	if node != nil {
		inOrderTraverse(node.left, printFunc)  // 先打印左子树
		printFunc(node.value)                  // 再打印根结点
		inOrderTraverse(node.right, printFunc) // 最后打印右子树
	}
}

//层级遍历
func (tree *ItemBinarySearchTree) LevelOrderTraverse(printFunc func(Item)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	levelOrderTraverse(tree.root, printFunc)
}
func levelOrderTraverse(node *Node, printFunc func(Item)) {
	if node == nil {
		return
	}
	var q NodeItemQueue
	q.New()
	q.Enqueue(node)
	for !q.IsEmpty() {
		if nodeTemp := q.Dequeue(); nodeTemp != nil {
			printFunc(nodeTemp.value)
			if nodeTemp.left != nil {
				q.Enqueue(nodeTemp.left)
			}
			if nodeTemp.right != nil {
				q.Enqueue(nodeTemp.right)
			}
		}
	}

}

//相等
func (tree *ItemBinarySearchTree) Equal(target *ItemBinarySearchTree) bool {
	if tree == nil || target == nil {
		return tree == nil && target == nil
	}
	return tree.root.Equal(target.root) //比较
}

//同构
//给定两棵树T1和T2.如果T1可以通过若干次左右孩子互换就变成T2，则我们称两棵树是同构的。
func (tree *ItemBinarySearchTree) Isomorphic(target *ItemBinarySearchTree) bool {
	tree.lock.RLock()
	target.lock.RLock()
	defer tree.lock.RUnlock()
	defer target.lock.RUnlock()
	var l1 ItemBinarySearchTree
	var l2 ItemBinarySearchTree
	var r1 ItemBinarySearchTree
	var r2 ItemBinarySearchTree

	if tree == nil || target == nil {
		return tree == nil && target == nil
	} else if tree.root == nil || target.root == nil {
		return tree.root == nil && target.root == nil
	} else if result := tree.root.SelfEqual(target.root); !result {
		return false
	} else if tree.root.left == nil && target.root.left == nil {
		r1.root = tree.root.right
		r2.root = target.root.right
		return r1.Isomorphic(&r2)
	} else if tree.root.left != nil && tree.root.left.SelfEqual(target.root.left) {
		l1.root = tree.root.left
		l2.root = target.root.left
		r1.root = tree.root.right
		r2.root = target.root.right
		return l1.Isomorphic(&l2) && r1.Isomorphic(&r2)
	} else if tree.root.left.SelfEqual(target.root.right) && tree.root.right.SelfEqual(target.root.left) {
		l1.root = tree.root.left
		l2.root = target.root.left
		r1.root = tree.root.right
		r2.root = target.root.right
		return l1.Isomorphic(&r2) && r1.Isomorphic(&l2)
	} else {
		return false
	}

}

// 打印树结构
func (tree *ItemBinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	if tree.root == nil {
		println("Tree is empty")
		return
	}
	stringify(tree.root, 0)
	println("----------------------------")
}
func stringify(node *Node, level int) {
	if node == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += "\t" // 根据节点的深度决定缩进长度
	}
	format += "----[ "
	level++
	/// 先递归打印右子树
	stringify(node.right, level)
	fmt.Printf(format+"%d\n", node.key)
	// 再递归打印左子树
	stringify(node.left, level)
}
