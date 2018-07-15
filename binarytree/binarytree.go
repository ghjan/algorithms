package binarytree

import (
	"github.com/ghjan/algorithms/queue"
	"github.com/ghjan/algorithms/set"
	"github.com/kataras/iris/core/errors"
)

/*
二叉树和二叉树节点
*/

//Node 二叉树节点
type Node struct {
	Data  rune
	Left  int
	Right int
}

//PreOrderTraverse 前序遍历
func (node *Node) PreOrderTraverse(tree BinaryTree, operationFunc func(nodeMe *Node)) {
	if node != nil {
		operationFunc(node) // 先打印根结点
		if node.Left >= 0 {
			tree[node.Left].PreOrderTraverse(tree, operationFunc) // 再打印左子树
		}
		if node.Right >= 0 {
			tree[node.Right].PreOrderTraverse(tree, operationFunc) // 最后打印右子树
		}
	}
}

//InOrderTraverse 中序遍历
func (node *Node) InOrderTraverse(tree BinaryTree, operationFunc func(nodeMe *Node)) {
	if node != nil {
		if node.Left >= 0 {
			tree[node.Left].InOrderTraverse(tree, operationFunc) // 先打印左子树
		}
		operationFunc(node) // 再打印根结点
		if node.Right >= 0 {
			tree[node.Right].InOrderTraverse(tree, operationFunc) // 最后打印右子树
		}
	}
}

//PostOrderTraverse 后序遍历(操作函數）：左子树 -> 根节点 -> 右子树
func (node *Node) PostOrderTraverse(tree BinaryTree, operationFunc func(nodeMe *Node)) {
	if node != nil {
		if node.Left >= 0 {
			tree[node.Left].PostOrderTraverse(tree, operationFunc) // 先打印左子树
		}
		if node.Right >= 0 {
			tree[node.Right].PostOrderTraverse(tree, operationFunc) // 再打印右子树
		}
		operationFunc(node) // 最后打印根结点
	}
}

//LevelOrderTraverse 层级遍历
func (node *Node) LevelOrderTraverse(tree BinaryTree, operationFunc func(node *Node)) {
	if node == nil {
		return
	}
	var q queue.ItemQueue
	q.New()
	q.Enqueue(*node)
	for !q.IsEmpty() {
		if nodeTemp := q.Dequeue(); nodeTemp != nil {
			no := (*nodeTemp).(Node)
			operationFunc(&no)
			if no.Left >= 0 {
				q.Enqueue(tree[no.Left])
			}
			if no.Right >= 0 {
				q.Enqueue(tree[no.Right])
			}
		}
	}

}

//InsertCode 插入编码的节点
//返回false：重复插入
func (node *Node) InsertCode(tree BinaryTree, data rune, code string) (*Node, error) {
	current := node
	for index, c := range code {
		parent := current
		isLeft := c == '0'
		currentIndex := -1
		if isLeft {
			currentIndex = current.Left
		} else {
			currentIndex = current.Right
		}
		if currentIndex == -1 { //空，需要创建新节点
			if index == len(code)-1 { // 叶子
				current = &Node{data, -1, -1}
			} else { //非叶子
				current = &Node{rune(0), -1, -1}
			}
			tree[index]=*current
			if isLeft {
				parent.Left = index
			} else {
				parent.Right = index
			}
		} else {
			current = &tree[currentIndex]
			if index == len(code)-1 { //重复插入
				return current, errors.New("重复插入")
			} else {
				if current.Data != rune(0) {
					return nil, errors.New("歧义编码")
				}
			}
		}
	}
	return current, nil
}

func (node Node) IsLeaf() bool {
	return node.Left == -1 && node.Right == -1
}

//BinaryTree 二叉树
type BinaryTree []Node

//创建二叉树
func CreateBinaryTree(arr []int) BinaryTree {
	N := len(arr)
	d := make([]Node, N, N)

	for i := 0; i < N; i++ {
		d[i].Left = -1
		d[i].Right = -1
	}
	for i, ar := range arr {
		d[i].Data = rune(ar)
	}
	for i := 0; i < len(arr)/2; i++ {
		if i*2+1 < len(d) {
			d[i].Left = i*2 + 1
		}
		if i*2+2 < len(d) {
			d[i].Right = i*2 + 2
		}
	}
	return d
}

func InitNewTree(cap int) BinaryTree {
	var tree BinaryTree
	tree = make([]Node, cap, cap)
	for i := 0; i < cap; i++ {
		tree[i].Left = -1
		tree[i].Right = -1
	}
	return tree
}
func (tree BinaryTree) Insert(data rune, index, left, right int) {
	if index < 0 {
		return
	}
	tree[index] = Node{data, left, right}
}

//Root 二叉树的根节点
func (tree BinaryTree) RootIndex() int {
	rootIndex := -1
	notRootSet := set.ItemSet{}
	for _, node := range tree {
		if node.Left != -1 {
			notRootSet.Add(node.Left)
		}
		if node.Right != -1 {
			notRootSet.Add(node.Right)
		}
	}
	for index := 0; index < len(tree); index++ {
		if !notRootSet.Has(index) {
			rootIndex = index
			break
		}
	}
	return rootIndex
}

func (tree BinaryTree) Isomorphic(T2 BinaryTree, indexR1, indexR2 int) bool {
	if (indexR1 == -1) || (indexR2 == -1) { //只要有一个是nil,就检查是否都nil
		return (indexR1 == -1) && (indexR2 == -1)
	}
	R1 := tree[indexR1]
	R2 := T2[indexR2]
	if R1.Data != R2.Data { //节点本身不一样，就不算同构
		return false
	} else if (R1.Left == -1) && (R2.Left == -1) {
		//如果左边全部是nil，那么就只要检查是否右子树同构
		return tree.Isomorphic(T2, R1.Right, R2.Right)
	} else if (R1.Right == -1) && (R2.Right == -1) {
		//如果右边全部是nil，那么就只要检查是否右子树同构
		return tree.Isomorphic(T2, R1.Left, R2.Left)
	} else if R1.Left == -1 && R2.Right == -1 {
		//如果都有一边是nil，
		return tree.Isomorphic(T2, R1.Right, R2.Left)
	} else if R1.Right == -1 && R2.Left == -1 {
		//如果都有一边是nil，
		return tree.Isomorphic(T2, R1.Left, R2.Right)
	} else if ((R1.Left >= 0) && (R2.Left >= 0)) &&
		((tree[R1.Left].Data) == (T2[R2.Left].Data)) { //如果左边都不是nil，而且left本身相等，那么检查左子树同构和右子树同构
		return tree.Isomorphic(T2, R1.Left, R2.Left) && tree.Isomorphic(T2, R1.Right, R2.Right)
	} else if (R1.Left >= 0) && (R2.Left >= 0) && (R1.Right >= 0) && (R2.Right >= 0) && (tree[R1.Left].Data == T2[R2.Right].Data) &&
		(tree[R1.Right].Data == T2[R2.Left].Data) { //如果左子树和右子树元素相同，而且右子树和左子树元素相同，那么就检查左子树和右子树是否同构
		return (tree.Isomorphic(T2, R1.Left, R2.Right)) && (tree.Isomorphic(T2, R1.Right, R2.Left))
	} else { //其他情况都非同构
		return false
	}
}

//PreOrderTraverse 前序遍历
func (tree BinaryTree) PreOrderTraverse(operationFunc func(nodeMe *Node)) {
	rootIndex := tree.RootIndex()
	if rootIndex >= 0 {
		tree[rootIndex].PreOrderTraverse(tree, operationFunc)
	}
}

//InOrderTraverse 中序遍历
func (tree BinaryTree) InOrderTraverse(operationFunc func(nodeMe *Node)) {
	tree[tree.RootIndex()].InOrderTraverse(tree, operationFunc)
}

//PostOrderTraverse 后序遍历(操作函數）：左子树 -> 根节点 -> 右子树
func (tree BinaryTree) PostOrderTraverse(operationFunc func(nodeMe *Node)) {
	tree[tree.RootIndex()].PostOrderTraverse(tree, operationFunc)
}

//LevelOrderTraverse 层级遍历
func (tree BinaryTree) LevelOrderTraverse(operationFunc func(nodeMe *Node)) {
	if tree.RootIndex() >= 0 {
		tree[tree.RootIndex()].LevelOrderTraverse(tree, operationFunc)
	}
}

type SimpleNode struct {
	Data  rune
	Left  int
	Right int
}

func (simpleNode SimpleNode) IsLeaf() bool {
	return simpleNode.Left == -1 && simpleNode.Right == -1
}

type SimpleBinaryTree []SimpleNode

//层级遍历
func (tree SimpleBinaryTree) LevelOrderTraverse(rootNode int, operationFunc func(SimpleNode)) {
	var q queue.ItemQueue
	q.New()
	q.Enqueue(tree[rootNode])
	for !q.IsEmpty() {
		if nodeTemp := q.Dequeue(); nodeTemp != nil {
			nodeObj := (*nodeTemp).(SimpleNode)
			operationFunc(nodeObj)
			if nodeObj.Left != -1 {
				q.Enqueue(tree[nodeObj.Left])
			}
			if nodeObj.Right != -1 {
				q.Enqueue(tree[nodeObj.Right])
			}
		}
	}

}
