package binarytree

import "errors"

//Node 二叉树节点
type Node struct {
	data  rune
	code  string
	left  *Node
	right *Node
}

//BinaryTree 二叉树
type BinaryTree []Node

//创建二叉树
func CreateTree(arr []int) BinaryTree {
	d := make([]Node, 0)
	for i, ar := range arr {
		d = append(d, Node{})
		d[i].data = rune(ar)
	}
	for i := 0; i < len(arr)/2; i++ {
		if i*2+1 < len(d) {
			d[i].left = &d[i*2+1]
		}
		if i*2+2 < len(d) {
			d[i].right = &d[i*2+2]
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
			current = current.left
		} else {
			current = current.right
		}
		if current == nil { //空，需要创建新节点
			if index == len(code)-1 { // 叶子
				current = &Node{data, code, nil, nil}
			} else { //非叶子
				current = &Node{rune(0), "", nil, nil}
			}
			if isLeft {
				parent.left = current
			} else {
				parent.right = current
			}
		} else {
			if index == len(code)-1 { //重复插入
				return current, errors.New("重复插入")
			} else {
				if current.data != rune(0) {
					return nil, errors.New("歧义编码")
				}
			}
		}
	}
	return current, nil
}

//前序遍历
func PreOrderTraverse(node *Node, operationFunc func(nodeMe *Node)) {
	if node != nil {
		operationFunc(node)                         // 打印根结点
		PreOrderTraverse(node.left, operationFunc)  // 先打印左子树
		PreOrderTraverse(node.right, operationFunc) // 再打印右子树
	}
}

//中序遍历
func InOrderTraverse(node *Node, operationFunc func(nodeMe *Node)) {
	if node != nil {
		InOrderTraverse(node.left, operationFunc)  // 先打印左子树
		operationFunc(node)                        // 打印根结点
		InOrderTraverse(node.right, operationFunc) // 再打印右子树
	}
}

// 后序遍历2操作函數）：左子树 -> 根节点 -> 右子树
func PostOrderTraverse(node *Node, operationFunc func(nodeMe *Node)) {
	if node != nil {
		PostOrderTraverse(node.left, operationFunc)  // 先打印左子树
		PostOrderTraverse(node.right, operationFunc) // 再打印右子树
		operationFunc(node)                          // 最后打印根结点
	}
}

//层级遍历
func LevelOrderTraverse(node *Node, operationFunc func(*Node)) {
	if node == nil {
		return
	}
	var q NodeItemQueue
	q.New()
	q.Enqueue(node)
	for !q.IsEmpty() {
		if nodeTemp := q.Dequeue(); nodeTemp != nil {
			operationFunc(nodeTemp)
			if nodeTemp.left != nil {
				q.Enqueue(nodeTemp.left)
			}
			if nodeTemp.right != nil {
				q.Enqueue(nodeTemp.right)
			}
		}
	}

}
