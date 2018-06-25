package binarysearchtree

import (
	"github.com/cheekybits/genny/generic"
	"fmt"
	"sync"
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

// 向树中插入元素
func (tree *ItemBinarySearchTree) Insert(key int, value Item) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	newNode := &Node{key, value, nil, nil}
	// 初始化树
	if tree.root == nil {
		tree.root = newNode
	} else {
		// 在树中递归查找正确的位置并插入
		insertNode(tree.root, newNode)
	}
}
func insertNode(node, newNode *Node) {
	// 插入到左子树
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			// 递归查找左边插入
			insertNode(node.left, newNode)
		}
	} else {
		// 插入到右子树
		if node.right == nil {
			node.right = newNode
		} else {
			// 递归查找右边插入
			insertNode(node.right, newNode)
		}
	}
}

// 搜索序号
// 返回true说明找到了
func (tree *ItemBinarySearchTree) Search(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return search(tree.root, key)
}
func search(node *Node, key int) bool {
	if node == nil {
		return false
	}
	// 向左搜索更小的值
	if key < node.key {
		return search(node.left, key)
	}
	// 向右搜索更大的值
	if key > node.key {
		return search(node.right, key)
	}
	return true // key == node.key
}

// 删除节点
//返回被删除的结点
func (tree *ItemBinarySearchTree) Remove(key int) (removedNode, newNode *Node) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	return remove(nil, tree.root, key, true)
}

// 递归删除节点
func remove(parent, node *Node, key int, is_left bool) (removedNode, newNode *Node) {
	// 要删除的节点不存在
	if node == nil {
		return nil, nil
	}

	var new_node *Node = nil
	// 寻找节点
	// 要删除的节点在左侧
	if key < node.key {
		node.left, new_node = remove(node, node.left, key, true)
		return node, new_node
	}
	// 要删除的节点在右侧
	if key > node.key {
		node.right, new_node = remove(node, node.right, key, false)
		return node, new_node
	}

	// 判断节点类型
	// 要删除的节点是叶子节点，直接删除
	// if key == node.key {
	if node.left == nil && node.right == nil {
		process_remove(parent, nil, is_left)
		return node, nil
	}

	// 要删除的节点只有一个节点，删除自身
	if node.left == nil {
		process_remove(parent, node.right, is_left)
		return node, node.right
	}
	if node.right == nil {
		process_remove(parent, node.left, is_left)
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
	new_node.key, new_node.value = mostLeftNode.key, mostLeftNode.value
	process_remove(parent, new_node, is_left)
	_, new_node.right = remove(new_node, new_node.right, new_node.key, false)
	return node, new_node
}
func process_remove(parent, node *Node, is_left bool) {
	if parent != nil {
		if is_left {
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

// 先序遍历：根节点 -> 左子树 -> 右子树
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

// 中序遍历：左子树 -> 根节点 -> 右子树
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

// 后续遍历：左子树 -> 右子树 -> 根结点
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
