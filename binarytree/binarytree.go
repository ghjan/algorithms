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
	Left  *Node
	Right *Node
}

//PreOrderTraverse 前序遍历
func (node *Node) PreOrderTraverse(operationFunc func(nodeMe *Node)) {
	if node != nil {
		operationFunc(node)                        // 先打印根结点
		node.Left.PreOrderTraverse(operationFunc)  // 再打印左子树
		node.Right.PreOrderTraverse(operationFunc) // 最后打印右子树
	}
}

//InOrderTraverse 中序遍历
func (node *Node) InOrderTraverse(operationFunc func(nodeMe *Node)) {
	if node != nil {
		node.Left.InOrderTraverse(operationFunc)  // 先打印左子树
		operationFunc(node)                       // 再打印根结点
		node.Right.InOrderTraverse(operationFunc) // 最后打印右子树
	}
}

//PostOrderTraverse 后序遍历(操作函數）：左子树 -> 根节点 -> 右子树
func (node *Node) PostOrderTraverse(operationFunc func(nodeMe *Node)) {
	if node != nil {
		node.Left.PostOrderTraverse(operationFunc)  // 先打印左子树
		node.Right.PostOrderTraverse(operationFunc) // 再打印右子树
		operationFunc(node)                         // 最后打印根结点
	}
}

//LevelOrderTraverse 层级遍历
func (node *Node) LevelOrderTraverse(operationFunc func(*Node)) {
	if node == nil {
		return
	}
	var q NodeItemQueue
	q.New()
	q.Enqueue(node)
	for !q.IsEmpty() {
		if nodeTemp := q.Dequeue(); nodeTemp != nil {
			operationFunc(nodeTemp)
			if nodeTemp.Left != nil {
				q.Enqueue(nodeTemp.Left)
			}
			if nodeTemp.Right != nil {
				q.Enqueue(nodeTemp.Right)
			}
		}
	}

}

//InsertCode 插入编码的节点
// 返回false：重复插入
func (node *Node) InsertCode(data rune, code string) (*Node, error) {
	current := node
	for index, c := range code {
		parent := current
		isLeft := c == '0'
		if isLeft {
			current = current.Left
		} else {
			current = current.Right
		}
		if current == nil { //空，需要创建新节点
			if index == len(code)-1 { // 叶子
				current = &Node{data, nil, nil}
			} else { //非叶子
				current = &Node{rune(0), nil, nil}
			}
			if isLeft {
				parent.Left = current
			} else {
				parent.Right = current
			}
		} else {
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
	return node.Left == nil && node.Right == nil
}

//BinaryTree 二叉树
type BinaryTree []Node

//创建二叉树
func CreateBinaryTree(arr []int) BinaryTree {
	d := make([]Node, 0)
	for i, ar := range arr {
		d = append(d, Node{})
		d[i].Data = rune(ar)
	}
	for i := 0; i < len(arr)/2; i++ {
		if i*2+1 < len(d) {
			d[i].Left = &d[i*2+1]
		}
		if i*2+2 < len(d) {
			d[i].Right = &d[i*2+2]
		}
	}
	return d
}

func CreateNewTree(cap int) BinaryTree {
	var tree BinaryTree
	tree = make([]Node, cap)
	return tree
}
func (tree BinaryTree) Insert(data rune, index, left, right int) {
	if index < 0 {
		return
	}
	var leftNode *Node = nil
	var rightNode *Node = nil
	if left > 0 {
		leftNode = &tree[left]
	}
	if right > 0 {
		rightNode = &tree[right]
	}
	tree[index] = Node{data, leftNode, rightNode}
}

//Root 二叉树的根节点
func (tree BinaryTree) RootIndex() int {
	rootIndex := -1
	notRootSet := set.ItemSet{}
	for _, node := range tree {
		if node.Left != nil {
			notRootSet.Add(*node.Left)
		}
		if node.Right != nil {
			notRootSet.Add(*node.Right)
		}
	}
	for index, node := range tree {
		if !notRootSet.Has(node) {
			rootIndex = index
			break
		}
	}
	return rootIndex
}

func (tree BinaryTree) Isomorphic(T2 BinaryTree, R1, R2 *Node) bool {
	if (R1 == nil) || (R2 == nil) { //只要有一个是nil,就检查是否都nil
		return (R1 == nil) && (R2 == nil)
	} else if R1.Data != R2.Data { //节点本身不一样，就不算同构
		return false
	} else if (R1.Left == nil) && (R2.Left == nil) {
		//如果左边全部是nil，那么就只要检查是否右子树同构
		return tree.Isomorphic(T2, R1.Right, R2.Right)
	} else if ((R1.Left != nil) && (R2.Left != nil)) &&
		((R1.Left.Data) == (R2.Left.Data)) { //如果左边都不是nil，而且left本身相等，那么检查左子树同构和右子树同构
		return tree.Isomorphic(T2, R1.Left, R2.Left) && tree.Isomorphic(T2, R1.Right, R2.Right)
	} else if (R1.Left.Data == R2.Right.Data) &&
		(R1.Right.Data == R2.Left.Data) { //如果左子树和右子树元素相同，而且右子树和左子树元素相同，那么就检查左子树和右子树是否同构
		return (tree.Isomorphic(T2, R1.Left, R2.Right)) && (tree.Isomorphic(T2, R1.Right, R2.Left))
	} else { //其他情况都非同构
		return false
	}
}

//PreOrderTraverse 前序遍历
func (tree BinaryTree) PreOrderTraverse(operationFunc func(nodeMe *Node)) {
	tree[tree.RootIndex()].PreOrderTraverse(operationFunc)
}

//InOrderTraverse 中序遍历
func (tree BinaryTree) InOrderTraverse(operationFunc func(nodeMe *Node)) {
	tree[tree.RootIndex()].InOrderTraverse(operationFunc)
}

//PostOrderTraverse 后序遍历(操作函數）：左子树 -> 根节点 -> 右子树
func (tree BinaryTree) PostOrderTraverse(operationFunc func(nodeMe *Node)) {
	tree[tree.RootIndex()].PostOrderTraverse(operationFunc)
}

//LevelOrderTraverse 层级遍历
func (tree BinaryTree) LevelOrderTraverse(operationFunc func(nodeMe *Node)) {
	tree[tree.RootIndex()].LevelOrderTraverse(operationFunc)
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
