package huffman

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
	var q queue.ItemQueue
	q.New()
	q.Enqueue(*node)
	for !q.IsEmpty() {
		if nodeTemp := q.Dequeue(); nodeTemp != nil {
			nt := (*nodeTemp).(Node)
			operationFunc(&nt)
			if nt.Left != nil {
				q.Enqueue(*nt.Left)
			}
			if nt.Right != nil {
				q.Enqueue(*nt.Right)
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
