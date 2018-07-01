package binarytree

import (
	"errors"

	"github.com/ghjan/algorithms/queue"
)

//Node 二叉树节点
type Node struct {
	Data  rune
	Left  *Node
	Right *Node
}
type SimpleNode struct {
	Data  rune
	Left  int
	Right int
}

func (node Node) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func (simpleNode SimpleNode) IsLeaf() bool {
	return simpleNode.Left == -1 && simpleNode.Right == -1
}

//BinaryTree 二叉树
type BinaryTree []Node
type SimpleBinaryTree []SimpleNode

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

//InsertCode 插入编码的节点
// 返回false：重复插入
func InsertCode(node *Node, data rune, code string) (*Node, error) {
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

//前序遍历
func (node *Node) PreOrderTraverse(operationFunc func(nodeMe *Node)) {
	if node != nil {
		operationFunc(node)                        // 先打印根结点
		node.Left.PreOrderTraverse(operationFunc)  // 再打印左子树
		node.Right.PreOrderTraverse(operationFunc) // 最后打印右子树
	}
}

//中序遍历
func (node *Node) InOrderTraverse(operationFunc func(nodeMe *Node)) {
	if node != nil {
		node.Left.InOrderTraverse(operationFunc)  // 先打印左子树
		operationFunc(node)                       // 再打印根结点
		node.Right.InOrderTraverse(operationFunc) // 最后打印右子树
	}
}

// 后序遍历2操作函數）：左子树 -> 根节点 -> 右子树
func (node *Node) PostOrderTraverse(operationFunc func(nodeMe *Node)) {
	if node != nil {
		node.Left.PostOrderTraverse(operationFunc)  // 先打印左子树
		node.Right.PostOrderTraverse(operationFunc) // 再打印右子树
		operationFunc(node)                         // 最后打印根结点
	}
}

//层级遍历
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
