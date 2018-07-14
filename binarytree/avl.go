package binarytree

import (
	"fmt"

	"github.com/ghjan/algorithms/queue"
	"github.com/kataras/iris/core/errors"
)

/*
AVL:自平衡二叉查找树
 */
//AvlTreeNode avl树节点
type AvlTreeNode struct {
	Key    int
	height int //高度：当前结点到叶子节点的距离 只有一个root节点高度为0
	depth  int //深度:当前节点到root的距离， root本身的深度是0
	left   *AvlTreeNode
	right  *AvlTreeNode
}

//NewAVLTreeNode 产生一个新节点
func NewAVLTreeNode(key int) *AvlTreeNode {
	return &AvlTreeNode{Key: key}
}

//Height avl树节点的高度
func (avlNode *AvlTreeNode) Height() int {
	if avlNode == nil {
		return -1
	} else {
		return avlNode.height
	}
}

//Depth avl树节点的深度
func (avlNode *AvlTreeNode) Depth() int {
	if avlNode == nil {
		return -1
	} else {
		return avlNode.depth
	}
}

//ProcessDepth 处理avlNode下面所有节点的深度  root.ProcessDepth(0)
func (avlNode *AvlTreeNode) ProcessDepth(depth int) {
	if avlNode == nil {
		return
	}
	avlNode.depth = depth
	if avlNode.left != nil {
		avlNode.left.ProcessDepth(depth + 1)
	}
	if avlNode.right != nil {
		avlNode.right.ProcessDepth(depth + 1)
	}
}

//层级遍历
func (avlNode *AvlTreeNode) LevelOrderTraverse(operationFunc func(*AvlTreeNode)) {
	if avlNode == nil {
		return
	}
	var q queue.ItemQueue
	q.New()
	q.Enqueue(*avlNode)
	for !q.IsEmpty() {
		if nodeTemp := q.Dequeue(); nodeTemp != nil {
			nodeObj := (*nodeTemp).(AvlTreeNode)
			operationFunc(&nodeObj)
			if nodeObj.left != nil {
				q.Enqueue(*nodeObj.left)
			}
			if nodeObj.right != nil {
				q.Enqueue(*nodeObj.right)
			}
		}
	}

}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// look LL kl和k之间的变化 kl上来 k作为kl的right kl原先的right放到k的左边
func leftLeftRotation(k *AvlTreeNode) *AvlTreeNode {
	var kl *AvlTreeNode
	kl = k.left
	k.left = kl.right
	kl.right = k
	k.height = max(k.left.Height(), k.right.Height()) + 1
	kl.height = max(kl.left.Height(), k.height) + 1
	return kl
}

//look RR k和kr之间的变化， kr上来，k作为kr的left， kr原先的left放到k的右边
func rightRightRotation(k *AvlTreeNode) *AvlTreeNode {
	var kr *AvlTreeNode
	kr = k.right
	k.right = kr.left
	kr.left = k
	k.height = max(k.left.Height(), k.right.Height()) + 1
	kr.height = max(k.height, kr.right.Height()) + 1
	return kr
}

//LR 先对左边做rightright 后leftleft
func leftRightRotation(k *AvlTreeNode) *AvlTreeNode {
	k.left = rightRightRotation(k.left)
	return leftLeftRotation(k)
}

//RL 先对右边做leftleft 后rightright
func rightLeftRotation(k *AvlTreeNode) *AvlTreeNode {
	k.right = leftLeftRotation(k.right)
	return rightRightRotation(k)
}

//Insert insert a Key to avl
func (avlNode *AvlTreeNode) Insert(key int) (*AvlTreeNode, error) {
	var err error
	if avlNode == nil {
		avlNode = NewAVLTreeNode(key)
	} else if key < avlNode.Key {
		avlNode.left, err = avlNode.left.Insert(key)
		if err == nil && avlNode.left.Height()-avlNode.right.Height() == 2 {
			if key < avlNode.left.Key { //LL
				avlNode = leftLeftRotation(avlNode)
			} else { // LR
				avlNode = leftRightRotation(avlNode)
			}
		}
	} else if key > avlNode.Key {
		avlNode.right, err = avlNode.right.Insert(key)
		if err == nil && (avlNode.right.Height()-avlNode.left.Height()) == 2 {
			if key < avlNode.right.Key { // RL
				avlNode = rightLeftRotation(avlNode)
			} else {
				//fmt.Println("right right", Key)
				avlNode = rightRightRotation(avlNode)
			}
		}
	} else if key == avlNode.Key {
		return avlNode, errors.New(fmt.Sprintf("the Key %d has existed!", key))
	}
	//notice: update height(may be this insert no rotation, so you should update height)
	avlNode.height = max(avlNode.left.Height(), avlNode.right.Height()) + 1
	return avlNode, nil
}

//DisplayAsc display avl tree  Key by asc
func (avlNode *AvlTreeNode) DisplayAsc() []int {
	return appendValues([]int{}, avlNode)
}

//DisplayDesc display avl tree Key by desc
func (avlNode *AvlTreeNode) DisplayDesc() []int {
	return appendValues2([]int{}, avlNode)
}

func appendValues(values []int, avl *AvlTreeNode) []int {
	if avl != nil {
		values = appendValues(values, avl.left)
		values = append(values, avl.Key)
		values = appendValues(values, avl.right)
	}
	return values
}

func appendValues2(values []int, avl *AvlTreeNode) []int {
	if avl != nil {
		values = appendValues2(values, avl.right)
		values = append(values, avl.Key)
		values = appendValues2(values, avl.left)
	}
	return values
}
